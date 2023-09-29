# docker buildx create --use
# docker buildx build -t hungddoo/telebot:latest --platform=linux/amd64,linux/arm64 --push .
# Build stage
FROM --platform=$BUILDPLATFORM golang:1.17-alpine AS builder
WORKDIR /src
COPY go.mod go.sum .
RUN go mod download
COPY . .
ARG TARGETOS TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /app ./src

# Runtime stage
FROM scratch
COPY --from=builder /app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV BOT_ENV=prod
ENTRYPOINT ["/app"]
