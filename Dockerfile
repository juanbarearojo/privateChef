FROM bitnami/golang:latest AS builder

ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM bitnami/golang:latest AS final

COPY --from=builder /go/bin/task /usr/local/bin/task

RUN adduser -D -h /home/test test

RUN mkdir -p /app/.cache/go-build \
    && chmod -R 777 /app/.cache/go-build 

USER test

ENV GOCACHE=/app/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
