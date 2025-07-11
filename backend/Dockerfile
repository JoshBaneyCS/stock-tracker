# Use official Golang image
FROM golang:1.25

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o server .

# Expose port for the Go server
EXPOSE 8080

# Run the compiled server binary
CMD ["./server"]
