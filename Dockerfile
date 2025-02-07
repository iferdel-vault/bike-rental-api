# ---------------------------------------
# Base Stage
# ---------------------------------------
FROM golang:1.23-alpine AS base

WORKDIR /app

# Add dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bikerentalapi ./cmd/api

# ---------------------------------------
# Development Build
# ---------------------------------------
FROM golang:1.23-alpine AS dev

WORKDIR /app

COPY ./air.toml ./air.toml

# Install Air and other necessary tools
RUN apk add --no-cache gcc musl-dev && \
    go install github.com/air-verse/air@latest

CMD ["air", "-c", "./air.toml"]

# ---------------------------------------
# Staging Build
# ---------------------------------------
FROM alpine:3.16 AS staging

WORKDIR /app

# Copy the compiled binary from the base stage
COPY --from=base /app/api /app/api

CMD ["/bin/sh", "/app/api"]

