FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o bot .
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /app/bot /app/
CMD ["/app/bot"]