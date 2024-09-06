# golang-react-todo
A full stack simple application using react and go

## Tools used for the App:
- **Frontend**: ReactJS
- **Backend**: GOlang

## Steps to run:

1. CLone the repo
2. Execute these commands in the terminal:
    ```
    # Apply frontend deployment and service
    kubectl apply -f frontend-deployment.yaml
    kubectl apply -f frontend-service.yaml

    # Apply backend deployment and service
    kubectl apply -f backend-deployment.yaml
    kubectl apply -f backend-service.yaml

    # Apply MongoDB deployment and service
    kubectl apply -f mongodb-deployment.yaml
    kubectl apply -f mongodb-service.yaml

    ```
3. Get the external IP of the frontend
    ```
    kubectl get svc
    ```
