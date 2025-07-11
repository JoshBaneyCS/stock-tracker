FROM golang:1.20

# Install Python & dependencies
RUN apt-get update && \
    apt-get install -y python3 python3-pip python3-venv && \
    apt-get clean

WORKDIR /app

# Copy files
COPY . .

# Create virtual environment & install Python requirements
RUN python3 -m venv /app/venv && \
    . /app/venv/bin/activate && \
    /app/venv/bin/pip install --upgrade pip && \
    /app/venv/bin/pip install -r requirements.txt

# Tidy Go modules
RUN go mod tidy

# Build Go app
RUN go build -o server .

CMD ["./server"]
