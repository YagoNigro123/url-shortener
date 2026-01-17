FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY public ./public

EXPOSE 8080

CMD ["./main"]