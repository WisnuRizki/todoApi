# Use an official Golang runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install dependencies using go modules
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 3030

# Run the Go app when the container starts
CMD ["go","run","main.go"]
