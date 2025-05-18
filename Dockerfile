FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/app .
COPY config ./config
COPY .env .env

CMD ["./app"]