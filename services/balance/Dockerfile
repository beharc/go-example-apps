FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o balance ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/balance .

EXPOSE 8080

CMD ["./balance"]
