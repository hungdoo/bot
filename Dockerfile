# Build stage
FROM golang:latest AS builder
WORKDIR /src
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app ./src

# Runtime stage
FROM scratch
COPY --from=builder /app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV BOT_ENV=prod
ENTRYPOINT ["/app"]
