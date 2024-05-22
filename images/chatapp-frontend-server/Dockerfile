# Use the official Go image as a parent image
FROM golang:1.21.6

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the staticserver source code into the container
COPY staticserver.go .

# Copy the static frontend files into the container
# Notice that the COPY paths have been adjusted since we're building from the chatapp-frontend-server directory

COPY chatapp-frontend/html ./html
COPY chatapp-frontend/html/js ./html/js
COPY chatapp-frontend/html/static ./html/static


# Build the Go app
RUN go build -o staticserver

# Command to run the executable
CMD ["./staticserver", "-static-dir", "./html"]


