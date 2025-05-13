# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Compile the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

# Runtime stage
FROM gcr.io/distroless/static:nonroot

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Expose the port used by the application
EXPOSE 8080

# Command to start the application
CMD ["./main"]
