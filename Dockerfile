FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o quasar .

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/quasar ./quasar

EXPOSE 8080

ENTRYPOINT ["./quasar"]
