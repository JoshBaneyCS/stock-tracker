# Use Go 1.23 image
FROM golang:1.23

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum before code to cache dependencies
COPY  go.mod go.sum ./

# Tidy dependencies
RUN go mod tidy

# Copy backend source code
COPY backend/ .

# Build the server binary
RUN go build -o server .

# Expose the API port
EXPOSE 8080

# Run the backend server
CMD ["./server"]
