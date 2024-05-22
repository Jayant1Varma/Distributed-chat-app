# Use the official Go image as a parent image
FROM golang:1.21.6

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY chatapp.go .

# Build the Go app
RUN go build -o chatapp-backend

# Command to run the executable
CMD ["./chatapp-backend"]
