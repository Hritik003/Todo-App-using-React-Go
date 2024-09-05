package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang-react-todo-1/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // Used for retrieving environment variables from .env file
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers to allow requests from the frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")            // Allow the frontend origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Allowed HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Allow Content-Type header

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request down the chain
		next.ServeHTTP(w, r)
	})
}

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	collection = client.Database(dbName).Collection(collName)
	log.Printf("Collection instance created")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	payload, err := getAllTasks()
	if err != nil {
		http.Error(w, "Failed to return all tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDo
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id := params["id"]
	if err := taskComplete(id); err != nil {
		http.Error(w, "Failed to complete task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	if err := undoTask(params["id"]); err != nil {
		http.Error(w, "Failed to undo task", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	if err := deleteOneTask(params["id"]); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count, err := deleteAllTasks()
	if err != nil {
		http.Error(w, "Failed to delete all tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(count)
}

func taskComplete(task string) error {
	id, err := primitive.ObjectIDFromHex(task)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	return err
}

func getAllTasks() ([]primitive.M, error) {
	curr, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var results []primitive.M
	for curr.Next(context.Background()) {
		var result bson.M
		if err := curr.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := curr.Err(); err != nil {
		return nil, err
	}

	curr.Close(context.Background())
	return results, nil
}

func insertOneTask(task models.ToDo) {
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted task")
}

func undoTask(task string) error {
	id, err := primitive.ObjectIDFromHex(task)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	return err
}

func deleteOneTask(task string) error {
	id, err := primitive.ObjectIDFromHex(task)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(context.Background(), filter)
	return err
}

func deleteAllTasks() (int64, error) {
	result, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
