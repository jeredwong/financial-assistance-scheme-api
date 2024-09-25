# Start from the official Go image
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the Swagger YAML file
COPY --from=builder /app/api/swagger/swagger.yaml ./api/swagger/swagger.yaml

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]