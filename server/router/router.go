package router

import (
	"golang-react-todo-1/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},         // Allow only this origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Allow specific methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allow specific headers
		AllowCredentials: true,                                      // Enable credentials if needed (like cookies or Authorization headers)
	})

	router.Use(c.Handler)

	//mux is used to match the incoming requests to their respective handler functions.

	router.HandleFunc("/api/task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}
		// Handle actual API request here
	})

	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	return router
}
