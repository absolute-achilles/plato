ARG GO_VERSION="1.26.3"
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
  git \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

# Copy dependency files first (better layer caching)
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -o /app/binary ./cmd/api

# Final Stage
FROM scratch

# Copy SSL certs for HTTPS calls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /app/binary /server

# Required environment variables for deployment
ENV PORT=""

ENV GIN_MODE="release"
ENTRYPOINT [ "/server" ]
