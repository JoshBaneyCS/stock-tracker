# Use Golang 1.23 base image
FROM golang:1.23-bullseye

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY ./go.mod ./go.sum ./

# Download Go dependencies
RUN go mod tidy

# Copy the full source code into the container
COPY . .

# Build the Go binary
RUN go build -o server .

# Expose backend port
EXPOSE 8080

# Run the compiled binary
CMD ["./server"]
