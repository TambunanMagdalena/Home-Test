# Build stage
FROM golang:1.24.2-alpine AS build-stage
RUN apk add --no-cache git ca-certificates tzdata
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-binary

# Production stage  
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata bash postgresql-client
WORKDIR /
COPY --from=build-stage /go-binary /go-binary
EXPOSE 3005
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3005/health || exit 1
ENTRYPOINT ["/go-binary"]