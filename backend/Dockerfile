# Use official Golang image with Go 1.23
FROM golang:1.23-bullseye

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum before other files to leverage Docker caching
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod tidy

# Copy the rest of the backend source code
COPY . .

# Build the Go backend binary
RUN go build -o server .

# Expose port used by the Go server
EXPOSE 8080

# Set the command to run the compiled server binary
CMD ["./server"]
