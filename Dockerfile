# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Set working directory and install tools for compression and build
WORKDIR /app
RUN apk add --no-cache bash build-base git openssh upx && \
    rm -rf /var/cache/apk/*

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application with static linking and optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main . && \
    upx --best --lzma main

# Stage 2: Minimal Image
FROM scratch

# Set working directory
WORKDIR /app

# Copy the statically built and compressed binary
COPY --from=builder /app/main .

# Command to run the application
ENTRYPOINT ["./main"]
