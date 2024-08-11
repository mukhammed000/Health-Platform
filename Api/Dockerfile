FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY .env ./

RUN go build -o main .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env ./
COPY ./docs /app/docs
COPY --from=builder /app/config/model.conf ./config/
COPY --from=builder /app/config/policy.csv ./config/

EXPOSE 8080

CMD ["./main"]