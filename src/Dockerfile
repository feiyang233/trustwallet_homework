# Use the official Golang image with version 1.22 as the base image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY .. .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# Build the Go application
RUN go mod download && go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
