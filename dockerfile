FROM node:18 as frontend-build

WORKDIR /client

COPY /client/package.json ./
RUN npm i
COPY /client .
RUN npm run build

FROM golang:1.23 as backend-build

WORKDIR /server

COPY ./server/go.mod ./server/go.sum ./
RUN go mod download
COPY ./server .
RUN go mod tidy
RUN go mod verify
RUN go build -o /server/server .

FROM nginx:alpine

WORKDIR /app

COPY --from=frontend-build /client/build /usr/share/nginx/html
COPY --from=backend-build /server/server /app/server

EXPOSE 8080 80

CMD ["/app/server", "&", "nginx", "-g", "daemon off;"]


