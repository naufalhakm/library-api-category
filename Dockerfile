FROM golang:1.22-alpine

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go build -o library-api-category ./cmd/server

EXPOSE 8084

CMD ["./library-api-category"]
