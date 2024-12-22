FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN [ -f go.mod ] || go mod init babka_bot
#RUN go mod tidy || { echo "go mod tidy failed"; exit 1; }
COPY . .
RUN go build -o bot .
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /app/bot /app/
CMD ["/app/bot"]