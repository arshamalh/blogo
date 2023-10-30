FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -a -o blogo .

FROM node:18-alpine3.15 AS frontend
WORKDIR /ui
COPY ui .
RUN npm install
RUN npm run build

FROM alpine:3.16
COPY --from=builder /app/blogo .
COPY --from=frontend /ui/dist /ui
CMD ["./blogo"]
