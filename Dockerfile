# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o smartload ./cmd/server/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/smartload .

EXPOSE 8080

CMD ["./smartload"]
