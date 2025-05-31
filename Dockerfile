# Use the official Go image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Install Git and other dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 3000

# Run the binary
CMD ["./main"]
