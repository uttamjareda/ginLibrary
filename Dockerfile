# Use an official Golang runtime as a parent image
FROM golang:1.18-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the application's source code from the local directory to the workspace in the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
