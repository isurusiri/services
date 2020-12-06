# Start from go base image
FROM golang:alpine AS builder

# Add maintainer info
LABEL maintainer="Isuru Siriwardana <isurusiri@protonmail.com>"

# Install git for fetching dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. All Dependencies will be cached 
# if the go.mod and go.sum files are not changed
RUN go mod download

# Copy source from the current directory to the working directory
COPY . . 

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o foo .

# Start a new stage from the scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/foo .
# COPY --from=builder /app/.env .

# Expose the port 8080 to the outside
EXPOSE 8080

# Execute the binary
CMD ["./foo"]
