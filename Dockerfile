# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY main.go .

RUN go build -o app main.go

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /build/app .

EXPOSE 8080

CMD ["./app"]
