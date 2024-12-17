
FROM golang:1.23-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=1

RUN go build -o device-manager-api ./cmd/device-manager-api

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/device-manager-api /app

EXPOSE 8080

CMD ["/app/device-manager-api"]
