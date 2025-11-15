# Build stage - UPDATE KE 1.24.2
FROM golang:1.24.2-alpine AS build-stage

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-binary

# Production stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata bash postgresql-client

WORKDIR /

# Copy binary
COPY --from=build-stage /go-binary /go-binary

# Copy wait-for-it script
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 3005

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3005/health || exit 1

ENTRYPOINT ["/go-binary"]