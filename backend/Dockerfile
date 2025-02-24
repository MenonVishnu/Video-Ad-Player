# Use an official Golang image
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy and build the application
COPY . .
RUN go build -o backend-app

# Use a smaller base image for production
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/backend-app .

# Set the PORT environment variable (default to 8080 if not provided)
ENV PORT=8080

# Expose the application port
EXPOSE ${PORT}

# Run the app
CMD ["./backend-app"]