# Use Golang Alpine as the builder
FROM golang:alpine AS builder

# Install necessary dependencies (including SQLite support)
RUN apk add --no-cache gcc musl-dev libc-dev sqlite sqlite-dev

# Set working directory inside the container
WORKDIR /app

# Enable CGO for SQLite support
ENV CGO_ENABLED=1

# Copy Go dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Ensure necessary directories exist
RUN mkdir -p /app/data 

# Build the Go application
RUN go build -o backend-app

# Use a smaller Alpine image for production
FROM alpine:latest

WORKDIR /app

# Install SQLite in the final container
RUN apk add --no-cache sqlite sqlite-libs

# Ensure necessary directories exist
RUN mkdir -p /app/data

# Copy the compiled binary from the builder stage
COPY --from=builder /app/backend-app .

# Copy the dummydata.json if it exists 
COPY --from=builder /app/data/dummydata.json /app/data/dummydata.json

# Copy the environment file 
COPY --from=builder /app/.env .env

# Set correct permissions for the database directory
RUN chmod -R 777 /app/data 

# Expose application port
EXPOSE 8080

# Run the backend application
CMD ["./backend-app"]
