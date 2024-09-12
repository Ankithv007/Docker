# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o calculator .

# Stage 2: Create a minimal image with the Go binary
FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/calculator /calculator

ENTRYPOINT ["/calculator"]
