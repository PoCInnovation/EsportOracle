# Multi-stage build for Go backend
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Enable CGO off for static build
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o server ./

FROM alpine:3.20 AS runtime

WORKDIR /app
COPY --from=builder /app/server /app/server

# If you have a .env file at runtime, mount or copy it next to the binary
# Expose backend port
EXPOSE 8080

# Required env vars (documented):
# - PANDASCORE_API_TOKEN
# - ETHEREUM_RPC_URL
# - BET_CONTRACT_ADDRESS

ENTRYPOINT ["/app/server"]
