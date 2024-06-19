# Stage 1: Build the Go application
FROM golang:1.22.2 as build

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Install required packages
RUN apt-get update && apt-get install -y \
    libx11-dev \
    libxi-dev \
    libxcursor-dev \
    libxrandr-dev \
    libxinerama-dev \
    libxrender-dev \
    libgl1-mesa-dev \
    libglu1-mesa-dev \
    libxxf86vm-dev \
    build-essential

# Build the Go application
RUN go build -o weather-app

# Debugging step: List the contents of the /app directory
RUN ls -al /app

# Stage 2: Create a smaller image for the runtime
FROM alpine:latest

# Install certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built Go application from the build stage
COPY --from=build /app/weather-app .

RUN ls -al /root

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["/root/weather-app"]
