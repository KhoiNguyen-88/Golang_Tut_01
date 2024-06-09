# Use an official Golang runtime as a parent image
FROM golang:1.18-alpine

# Set the current working directory inside the container
WORKDIR /go/src/Golang_Tut_01

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Copy the .env file into the container
COPY .env .env

# Build the Go app
RUN go build -o /go/bin/main ./cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go/bin/main"]
