# Шаг 1: Используем официальный образ Go для сборки приложения
FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN if [ ! -f go.mod ]; then go mod init babka_bot && go mod tidy; fi
RUN go build -o bot .
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /app/bot /app/
ENV GIN_MODE=release
CMD ["/app/bot"]