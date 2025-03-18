FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o user ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user .

EXPOSE 8080

CMD ["./user"]
