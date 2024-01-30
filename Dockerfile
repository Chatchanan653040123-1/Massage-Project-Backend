# Use a builder stage to build the Go application
FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app
COPY . /app
RUN go build -o main main.go

# Build a smaller image for the final application
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY . /app/
EXPOSE 5000
CMD ["/app/main"]
