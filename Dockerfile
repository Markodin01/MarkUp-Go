# Start from a base Go image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the project files to the working directory
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port the application will run on
EXPOSE 8080

# Set the command to run the application
CMD ["go run main.go storage.go models.go handlers.go"]