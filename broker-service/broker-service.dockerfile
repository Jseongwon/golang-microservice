# Base Go image for building the app
FROM golang:1.18-alpine AS builder

# Install necessary tools
RUN apk add --no-cache git

# Set up working directory
WORKDIR /app

# Copy go.mod and go.sum first to cache dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o brokerApp ./cmd/api

# Create a lightweight image for running the application
FROM alpine:latest

# Set up working directory
WORKDIR /app

# Copy the built application binary
COPY --from=builder /app/brokerApp /app/brokerApp

# Expose the application port
EXPOSE 8080

# Set the entry point for the container
CMD [ "/app/brokerApp" ]