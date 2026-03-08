FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o uptime-scope .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/uptime-scope /app/uptime-scope

EXPOSE 8181

ENTRYPOINT ["/app/uptime-scope"]
