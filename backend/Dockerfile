FROM golang:1.24 AS builder

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /backend/main .

CMD ["./main"]