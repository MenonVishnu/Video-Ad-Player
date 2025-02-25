# Use Golang Alpine as the builder
FROM golang:alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache gcc musl-dev libc-dev sqlite sqlite-dev

# Set working directory inside container
WORKDIR /app

# Enable CGO for SQLite
ENV CGO_ENABLED=1

# Copy Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Ensure the database directory exists
RUN mkdir -p /app/data

# Build the Go application
RUN go build -o backend-app

# Use a smaller Alpine image for production
FROM alpine:latest

WORKDIR /app

# Install SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs

# Ensure the database directory exists in the final container
RUN mkdir -p /app/data

# Copy the compiled binary from the builder stage
COPY --from=builder /app/backend-app .

# Copy the .env file (if required)
COPY --from=builder /app/.env .

# Set correct permissions for the database directory
RUN chmod -R 777 /app/data

# Expose application port
EXPOSE 8080

# Run the backend application
CMD ["./backend-app"]
