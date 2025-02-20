FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -v -o ./bin/rinha ./cmd/main.go

FROM alpine:3.14.10

EXPOSE 8080

# Copy files from builder stage
COPY --from=builder /app/bin/rinha .

# Run binary
CMD ["/rinha"]