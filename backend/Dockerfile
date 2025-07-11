# Use Golang 1.23 base image
FROM golang:1.23-bullseye

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker cache
COPY go.mod go.sum ./

# Download Go modules
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o server .

# Expose the application port
EXPOSE 8080

# Start the server
CMD ["/app/server"]
