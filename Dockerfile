# Use a development image with Go installed
FROM golang:1.25-alpine

# Set the working directory inside the container
WORKDIR /app

# Install the Air utility using the correct, updated path
# This will fix the 'version constraints conflict' error
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Set the entry point to run Air. 
# Air will handle the initial compilation and subsequent reloads.
ENTRYPOINT ["air"]

# Expose any necessary ports (e.g., for a web server)
EXPOSE 8080
