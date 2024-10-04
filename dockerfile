# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first for dependency resolution
COPY go.mod go.sum ./

# Download all the dependencies. This will be cached if the files don't change
RUN go mod download

# Copy the rest of the application code
COPY . .

# Install godotenv for loading environment variables (if using .env files)
RUN go get github.com/joho/godotenv

# Build the Go app for production
RUN go build main.go

# Command to run the app
CMD ["./main"]
