# Builder
FROM golang:1.22-alpine AS builder
ADD . /src/notification
WORKDIR /src/notification
RUN [ -x "$(command -v task)" ] || go install github.com/go-task/task/v3/cmd/task@latest && \
    task -t Taskfile.notification.yaml update && \
    task -t Taskfile.notification.yaml format && \
    task -t Taskfile.notification.yaml test
RUN CGO_ENABLED=1 GOOS=linux go build -x -o /build/notification .

# Service
FROM alpine:latest
EXPOSE 8000
RUN chmod 777 /opt
COPY --from=builder /build/notification /usr/bin
USER nobody
ENTRYPOINT ["/usr/bin/notification"]
