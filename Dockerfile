# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy the entire project
COPY . .

# Build the Go app
ARG SERVICE_NAME
RUN cd services/${SERVICE_NAME} && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../../main ./cmd

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
