# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.17 AS builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Download and install any required dependencies.
RUN go mod download

# Build the binary.
RUN go build -o miniproject .

# Use a minimal alpine image to run the binary.
# https://hub.docker.com/_/alpine
FROM alpine:3.14
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/miniproject .

# Run the binary.
CMD ["./miniproject"]