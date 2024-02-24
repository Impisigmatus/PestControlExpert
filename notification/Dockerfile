# Builder
FROM golang:1.22-alpine AS builder
ADD . /src/notification
WORKDIR /src/notification
RUN [ -x "$(command -v task)" ] || go install github.com/go-task/task/v3/cmd/task@latest && \
    task notification:update && \
    task notification:format && \
    task notification:test
RUN CGO_ENABLED=1 GOOS=linux go build -x -o /build/notification .

# Service
FROM alpine:latest
EXPOSE 8000
RUN chmod 777 /opt
COPY --from=builder /build/notification /usr/bin
USER nobody
ENTRYPOINT ["/usr/bin/notification"]
