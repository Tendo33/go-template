FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/server

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /server /server

RUN addgroup -S app && adduser -S -G app app

USER app

EXPOSE 8080

CMD ["/server"]
