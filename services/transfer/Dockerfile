FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o transfer ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/transfer .

EXPOSE 8080

CMD ["./transfer"]
