# syntax=docker/dockerfile:1

# Build stage
FROM --platform=$BUILDPLATFORM golang:1.23 as builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -o pmcp main.go

# Final image
FROM alpine:3.22

WORKDIR /

COPY --from=builder /app/pmcp /pmcp

ENTRYPOINT ["/pmcp"]
