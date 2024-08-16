FROM golang:1.22.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env ./
RUN go build -o main

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8082

CMD ["./main"]
