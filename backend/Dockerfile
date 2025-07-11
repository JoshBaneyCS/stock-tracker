FROM golang:1.20

# Install Python
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    apt-get clean

# Set work directory
WORKDIR /app

# Copy Go backend
COPY . .

# Install Python deps
RUN pip3 install -r requirements.txt

# Build Go server
RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
