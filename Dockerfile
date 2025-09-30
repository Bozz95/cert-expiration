# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
# Using only standard libraries, no go.sum needed
COPY go.mod ./

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cert-expiration .

# Final stage
FROM scratch

# Copy CA certificates for HTTPS connections
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /app/cert-expiration /cert-expiration

# Set entrypoint
ENTRYPOINT ["/cert-expiration"]