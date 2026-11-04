FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

# BuildKit sets these from the target platform. A default here would replace
# that value and cross-compile the wrong architecture.
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -trimpath -o /server ./cmd/server

FROM alpine:3.22

RUN apk add --no-cache ca-certificates wget

WORKDIR /app

COPY --from=builder /server /server

RUN addgroup -S app && adduser -S -G app app

USER app

EXPOSE 8080

# PORT defaults to 8080. The shell form reads the runtime value.
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null "http://127.0.0.1:${PORT:-8080}/health" || exit 1

CMD ["/server"]
