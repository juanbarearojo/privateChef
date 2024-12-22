FROM ubuntu:latest AS builder

RUN apt-get update && \
    apt-get install -y git golang-go ca-certificates && \
    update-ca-certificates

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM ubuntu:latest AS final

RUN apt-get update && \
    apt-get install -y git golang-go ca-certificates && \
    update-ca-certificates

COPY --from=builder /root/go/bin/task /usr/local/bin/task

RUN useradd -m -d /home/test test

RUN mkdir -p /app/.cache/go-build \
    && chmod -R 777 /app/.cache/go-build

USER test

ENV GOCACHE=/app/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["task", "test"]