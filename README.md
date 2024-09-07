# golang-react-todo
A full stack simple application using react and go

## Tools used for the App:
- **Frontend**: ReactJS
- **Backend**: GOlang

## Steps to run:

1. CLone the repo
2. Execute these commands in the terminal:
    - Go to the ./client/Kubernetes and execute
        ```
        kubectl apply -f frontend.yaml 
        ```
    - Go to the ./server/Kubernetes and execute
        ```
        kubectl apply -f backend.yaml 
        ```
    - Go to the ./mongodb and execute
        ```
        kubectl apply -f mongodb.yaml 
        ```
3. Get the external IP of the frontend
    ```
    kubectl get svc
    ```
