# Stage 1: Build the Go binary
FROM golang:1.23.4-alpine AS builder

# Set working directory
WORKDIR /app

# Copy all codes since we are not using any external dependencies
COPY . .

# Build the binary
RUN go build -o sorter ./main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/sorter .

# Set default command
ENTRYPOINT ["./sorter"]
