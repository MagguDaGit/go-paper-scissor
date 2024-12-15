
# Stage 1: Build the application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the Go source files

COPY . ./

# Build the Go application (ensure the name matches the main Go file in your project)


RUN CGO_ENABLED=0 go build -o go-paper-scissor .


# Stage 2: Create a lightweight image with just the binary
FROM alpine:latest

WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/go-paper-scissor .

# Expose the port your app will listen on
EXPOSE 8080

# Command to run the binary
CMD ["./go-paper-scissor"]

