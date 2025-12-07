# syntax=docker/dockerfile:1

ARG VERSION_NUMBER=(unknown)

# Build stage
FROM --platform=$BUILDPLATFORM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION_NUMBER
ARG GIT_COMMIT
ARG BUILD_DATE
ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -o pmcp \
    -ldflags="-s -w -X 'github.com/yshngg/pmcp/internal/version.Number=${VERSION_NUMBER}' -X 'github.com/yshngg/pmcp/internal/version.GitCommit=${GIT_COMMIT}' -X 'github.com/yshngg/pmcp/internal/version.BuildDate=${BUILD_DATE}'" \
    .

# Final image
FROM alpine:3.23

WORKDIR /

COPY --from=builder /app/pmcp /pmcp

USER nobody

ENTRYPOINT ["/pmcp"]
