# syntax=docker/dockerfile:1

# Build stage
FROM --platform=$BUILDPLATFORM golang:1.24 AS builder

ARG TARGETOS
ARG TARGETARCH
ARG VERSION_NUMBER=0.0.0-dev
ARG GIT_COMMIT=unknown
ARG BUILD_DATE=unknown

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -o pmcp \
    -ldflags="-s -w -X 'github.com/yshngg/pmcp/pkg/version.Number=${VERSION_NUMBER}' -X 'github.com/yshngg/pmcp/pkg/version.GitCommit=${GIT_COMMIT}' -X 'github.com/yshngg/pmcp/pkg/version.BuildDate=${BUILD_DATE}'"

# Final image
FROM alpine:3.22

WORKDIR /

COPY --from=builder /app/pmcp /pmcp

ENTRYPOINT ["/pmcp"]
