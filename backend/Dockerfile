FROM golang:1.23

# Install Python, venv, and pip
RUN apt-get update && \
    apt-get install -y python3 python3-pip python3-venv && \
    apt-get clean

# Set working directory
WORKDIR /app

# Copy backend source code
COPY . .

# Create a virtual environment and install Python dependencies
RUN python3 -m venv /app/venv && \
    . /app/venv/bin/activate && \
    /app/venv/bin/pip install --no-cache-dir -r requirements.txt

# Build Go server
RUN go build -o server .

# Expose port
EXPOSE 8080

# Run the Go app and keep the Python venv active
CMD ["/app/server"]
