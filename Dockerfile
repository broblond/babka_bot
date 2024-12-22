FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bot ./main.go
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /app/bot /app/
CMD ["/app/bot"]