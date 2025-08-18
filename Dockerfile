FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /usr/local/bin/golearn ./cmd/golearn

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/golearn /usr/local/bin/golearn
WORKDIR /workspace
ENTRYPOINT ["golearn"]


