FROM okteto/golang:1.23 AS builder

RUN apt-get install git

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM okteto/golang:1.23 AS final

COPY --from=builder /go/bin/task /usr/local/bin/task

RUN adduser -D -h /home/test test

RUN mkdir -p /app/.cache/go-build \
    && chmod -R 777 /app/.cache/go-build 

USER test

ENV GOCACHE=/app/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
