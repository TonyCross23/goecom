# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Install migrate CLI inside builder
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build app binary
RUN go build -o bin/ecom cmd/main.go


FROM debian:bookworm-slim

WORKDIR /root/

# Copy migrate binary from builder
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy built app binary
COPY --from=builder /app/bin/ecom ./ecom

# Copy migrations
COPY cmd/migrate/migrations ./migrations

EXPOSE 8080

CMD ["./ecom"]