FROM golang:1.23.4-alpine3.21

WORKDIR /app

# Install necessary build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port your application runs on
EXPOSE 8080

CMD ["./main"]