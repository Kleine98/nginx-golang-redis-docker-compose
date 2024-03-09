# Use the official Golang image as a base
FROM golang:alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Create a new lightweight image containing only the built executable
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/app .

# Expose port 3000
EXPOSE 3000

# Command to run the Go application
CMD ["./app"]
