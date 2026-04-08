FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /go-learning-api ./cmd/server

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /go-learning-api /usr/local/bin/go-learning-api

EXPOSE 8080

CMD ["go-learning-api"]
