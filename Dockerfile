FROM golang:1.22.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start fresh from a smaller image
FROM alpine:3.17.1

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["/app/main"]