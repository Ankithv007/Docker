# Use the official Go image
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o calculator .

# Set the entrypoint to the executable
ENTRYPOINT ["./calculator"]
