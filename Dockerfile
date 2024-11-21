# Use the official Go image as the builder
FROM golang:1.23.3-alpine as builder

# Set the working directory
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the Go source code and .env file into the container
COPY . .

# Build the Go application
RUN go build -tags netgo -ldflags '-s -w' -o api ./_cmd

# Create the final image with a smaller base (Alpine)
FROM alpine:latest

# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Copy the compiled binary and .env file from the builder stage
COPY --from=builder /app/api /api
COPY .env /app/.env

# Expose the port the app will run on
EXPOSE 8082

# Command to run the application
CMD ["/api"]
