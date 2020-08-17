# Use the latest Go 1.X image for building
FROM golang:1 as builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create app directory
WORKDIR /usr/src/appbuild

# Copy the source files
COPY src src
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

# Download all the dependencies and build
RUN go get ./... && \
    go build -o ./bin/handler ./src

# Use alpine image for the actual deployment
FROM alpine:latest

# Create app directory
WORKDIR /usr/src/app

# Copy the app
COPY --from=builder /usr/src/appbuild/bin/handler bin/

# Bind the app port
EXPOSE 7200

# Start the app
ENTRYPOINT bin/handler
