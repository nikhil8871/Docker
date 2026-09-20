# Stage 1: Build the Go application
FROM golang:1.22-alpine AS build

WORKDIR /app

# Install git for any private or git dependencies
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /employee-management-system .

# Stage 2: Minimal, secure runtime
FROM alpine:3.19

WORKDIR /app

# Install security certificates, tzdata, and bash for wait-for-it script
RUN apk --no-cache add ca-certificates tzdata bash

# Create non-privileged user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy binary and startup script
COPY --from=build /employee-management-system /app/employee-management-system
COPY wait-for-it.sh /app/wait-for-it.sh

# Set permissions and ownership
RUN chmod +x /app/wait-for-it.sh && chown -R appuser:appgroup /app

# Enforce least-privilege non-root execution
USER appuser

EXPOSE 8080

CMD ["/app/employee-management-system"]
