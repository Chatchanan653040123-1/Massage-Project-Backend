# Use a builder stage to build the Go application
FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app
COPY . /app
RUN go build -o main main.go

# Build a smaller image for the final application
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY configs/ /app/configs/
COPY errors/ /app/errors/
COPY handlers/ /app/handlers/
COPY logs/ /app/logs/
COPY repositories/ /app/repositories/
COPY services/ /app/services/
COPY databases/ /app/databases/
EXPOSE 8080
CMD ["/app/main"]
