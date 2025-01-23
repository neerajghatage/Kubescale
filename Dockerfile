# Use an official Golang runtime as the base image
FROM golang:1.18-alpine

# Set the working directory in the container
WORKDIR /app

# Install necessary packages
RUN apk update && apk add --no-cache gcc musl-dev

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Specify the command to run your application
CMD ["./main"]
