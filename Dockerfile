FROM golang:1.21.6-alpine AS build-stage

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-binary

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata bash postgresql-client

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /

COPY --from=build-stage /go-binary /go-binary
COPY --from=build-stage /app/.env /.env

COPY --from=build-stage /app/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

RUN chown -R appuser:appgroup /go-binary /wait-for-it.sh

USER appuser

EXPOSE 3005

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3005/health || exit 1

ENTRYPOINT ["/go-binary"]