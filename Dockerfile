FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o bot .

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/bot /app/

CMD ["/app/bot"]
