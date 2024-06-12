# syntax=docker/dockerfile:1

# Builder stage
FROM golang:1.22 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY ./ ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-ascii-art-web

# Final stage
FROM alpine:latest

# Install the certificates to avoid "x509: certificate signed by unknown authority" errors
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /docker-ascii-art-web /docker-ascii-art-web

# Run the binary
CMD ["/docker-ascii-art-web"]

# Expose the port the application runs on
EXPOSE 8080
