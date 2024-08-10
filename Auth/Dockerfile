FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env ./

RUN go build -o main ./cmd

FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env ./

EXPOSE 8081

CMD ["./main"]
