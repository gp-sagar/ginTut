# Use the official golang image as the base image
FROM golang:1.17-alpine

# Set the Current Working Directory inside the container
WORKDIR /ginTut

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Copy the .env file to the container (if needed)
COPY .env .env

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
