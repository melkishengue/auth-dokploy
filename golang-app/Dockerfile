# golang-app/Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init example.com/app
RUN CGO_ENABLED=0 go build -o /app/main

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/main
EXPOSE 8081
CMD ["/app/main"]
