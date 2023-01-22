# Builder
FROM golang:1.19-alpine AS builder
ADD . /src/notification
WORKDIR /src/notification
RUN CGO_ENABLED=1 GOOS=linux go build -x -o /build/notification .

# Service
FROM alpine:latest
EXPOSE 8000
COPY --from=builder /build/notification /usr/bin
USER nobody
ENTRYPOINT ["/usr/bin/notification"]
