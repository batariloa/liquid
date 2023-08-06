# Use an official Golang runtime as a base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files to the working directory
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose the port on which your Go application listens (change the port if needed)
EXPOSE 8080

# Command to run the Go application
CMD ["./myapp"]

