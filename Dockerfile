FROM debian:bookworm-slim AS builder

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        wget \
        git && \
    update-ca-certificates

ENV GO_VERSION=1.23.2
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM debian:bookworm-slim

LABEL org.opencontainers.image.authors="juanbarearojo0@correo.ugr.es"

COPY --from=builder /usr/local/go /usr/local/go

ENV PATH=/usr/local/go/bin:$PATH

COPY --from=builder /go/bin/task /usr/local/bin/task

RUN useradd -m test && \
    mkdir -p /app/.cache/go-build && \
    chmod -R 777 /app/.cache/go-build

USER test

ENV GOCACHE=/app/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
