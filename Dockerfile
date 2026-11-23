# Multi-stage Docker build for CertiScholar Notifications
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download || true
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /certischolar-notifications ./cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /certischolar-notifications .
EXPOSE 8081
CMD ["./certischolar-notifications"]
