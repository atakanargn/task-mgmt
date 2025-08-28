# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app .

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/app .
# Uygulamanın çalışacağı port
EXPOSE 8080
# Konteyner başladığında çalışacak komut
CMD ["./app"]
