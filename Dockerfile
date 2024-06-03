#build
FROM golang:1.22-alpine as build

WORKDIR /app

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux  go build -ldflags="-w -s" -o rate-limiter cmd/main.go

EXPOSE 8080

CMD ["./rate-limiter"]