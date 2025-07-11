# Use Go base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /go/src/backend

# Copy everything from backend/ into the container
COPY . .

# Download Go module dependencies
RUN go mod tidy

# Build the Go backend binary
RUN go build -o server .

# Expose backend service port (adjust as needed)
EXPOSE 8080

# Run the compiled binary
CMD ["./server"]
