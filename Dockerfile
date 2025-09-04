FROM golang:1.25.1-alpine

# Install CompileDaemon for live reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Expose the application port
EXPOSE 8080

# Use CompileDaemon to rebuild and restart on code changes
ENTRYPOINT ["CompileDaemon", "--build=go build -o main ./cmd", "--command=./main"]

# Build stage
# FROM golang:1.24-alpine AS builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src

# Runtime stage
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY .env .
# EXPOSE 8080
# CMD ["./main"]
