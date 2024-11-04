FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o qlub .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/qlub /usr/local/bin/qlub

COPY subdomains.json .

CMD ["/usr/local/bin/qlub", "--config", "/app/subdomains.json", "--watch"]
