# Use the Go version 1.23.3 image
FROM golang:1.23.3

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download and verify dependencies
RUN go mod tidy

# Copy the remaining application files
COPY . .

# Build the Go binary
RUN go build -o main .

# Set the command to run the application
CMD ["./main"]
