# BUild statge
FROM golang:alpine AS builder

# Set necessary environment variables needed for our image
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
RUN go build -o main ./main.go

# Run stage
FROM alpine
WORKDIR /build
COPY --from=builder /build/main .

COPY scripts/entrypoint.sh .
COPY scripts/swagger.sh .
COPY test.db .

# Export necessary port
EXPOSE 8080

ENTRYPOINT ["/build/entrypoint.sh"]