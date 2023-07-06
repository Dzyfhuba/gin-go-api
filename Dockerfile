# Base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

RUN go install github.com/swaggo/swag/cmd/swag

# Download Go module dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Expose the necessary port
EXPOSE 3333

# Entrypoint command to run Go with bin/air
CMD ["bin/air"]
