FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN go mod init babka_bot || true && go mod tidy
RUN go build -o bot ./main.go
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /app/bot /app/
ENV GIN_MODE=release
CMD ["/app/bot"]