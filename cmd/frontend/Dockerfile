# Import GO_VERSION
ARG GO_VERSION

# Use the official Golang image as a base
FROM golang:${GO_VERSION} AS builder

# Set the working directory
WORKDIR go/src/app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Import service_path
ARG service_path

# Copy the source code
COPY ./${service_path} ./${service_path}
COPY ./pkg ./pkg

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./${service_path}

# Final stage
FROM scratch

# Copy binary from build stage
COPY --from=builder /go/bin/app /go/bin/app

# Expose the service port
EXPOSE 8080

# Set the entrypoint to run the service
ENTRYPOINT ["/go/bin/app"]