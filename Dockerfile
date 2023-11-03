# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose a port (if your application listens on a specific port)
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]
