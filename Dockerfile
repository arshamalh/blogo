# Stage 1: Build Go and Swagger
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -a -o blogo .

# Stage 2: Build Frontend
FROM node:18-alpine3.15 AS frontend
WORKDIR /ui
COPY ui .
RUN npm install
RUN npm run build

# Stage 3: Create the final image
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/blogo .
COPY --from=frontend /ui/dist /ui
COPY swagger.json /app/swagger.json
CMD ["./blogo", "serve", "--host", "0.0.0.0", "--port", "8080"]
