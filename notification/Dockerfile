# Builder
FROM golang:1.19-alpine AS builder
RUN go install github.com/go-task/task/v3/cmd/task@latest
ADD ../Taskfile.yaml /src
ADD . /src/notification
WORKDIR /src/notification
RUN task notification:autogen && \
    CGO_ENABLED=1 GOOS=linux go build -x -o /build/notification .

# Service
FROM alpine:latest
EXPOSE 8000
COPY --from=builder /build/notification /usr/bin
USER nobody
ENTRYPOINT ["/usr/bin/notification"]
