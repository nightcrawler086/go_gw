# Start with a minimal base image containing the Go runtime
FROM golang:1.20-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o go_gw

# Start with a fresh, minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/go_gw .

# Expose the port that the application listens on
EXPOSE 8080

# Define the command to run the application
CMD ["./go_gw"]

