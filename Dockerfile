FROM debian:latest AS builder

RUN apt-get update && \
    apt-get install -y git wget tar 

ENV GO_VERSION=1.23.2

RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM debian:latest AS final

COPY --from=builder /usr/local/go /usr/local/go

ENV PATH=$PATH:/usr/local/go/bin

COPY --from=builder /root/go/bin/task /usr/local/bin/task

RUN useradd -m -d /home/test test

RUN mkdir -p /app/.cache/go-build /app/test && \
    chmod -R 777 /app/.cache/go-build

USER test

ENV GOCACHE=/app/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
