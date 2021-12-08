FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o dist/apps/lukso-manager apps/lukso-manager/main.go

# Copy binary from build to main folder
RUN cp dist/apps/lukso-manager .

# Build a small image
FROM golang:1.17.3
RUN apt-get update && apt-get install bash

COPY --from=builder /build/lukso-manager /

# Command to run
ENTRYPOINT ["/lukso-manager"]