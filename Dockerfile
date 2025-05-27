# ── Build stage ───────────────────────────────────────────────────────────────
FROM golang:1.20-alpine AS builder

# Ensure modules are cached
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

# Copy in source, then compile statically for Linux
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

# ── Final stage ───────────────────────────────────────────────────────────────
FROM scratch

# Drop in the compiled binary and, if needed, any certs or static assets
COPY --from=builder /app/server /server

# (Optional) If your app needs CA certs, uncomment these lines:
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 6379

# Run as non-root (optional but recommended)
USER 65532:65532

ENTRYPOINT ["/server"]
