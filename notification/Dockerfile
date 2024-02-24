# Builder
FROM golang:1.19-alpine AS builder
ADD . /src/notification
WORKDIR /src/notification
RUN [ -x "$(command -v oapi-codegen)" ] || go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
    mkdir -p autogen && \
    sed -e 's/^/    /' README.md > autogen/README.md && \
    sed -e '/    %README.md%/{' -e "r autogen/README.md" -e 'd' -e '}' oapi3.yaml > autogen/oapi3.yaml && \
    oapi-codegen -o autogen/server.go -old-config-style -package autogen -generate chi-server autogen/oapi3.yaml && \
    oapi-codegen -o autogen/types.go  -old-config-style -package autogen -generate types      autogen/oapi3.yaml && \
    oapi-codegen -o autogen/client.go -old-config-style -package autogen -generate client     autogen/oapi3.yaml
RUN ls -R
RUN CGO_ENABLED=1 GOOS=linux go build -x -o /build/notification .

# Service
FROM alpine:latest
EXPOSE 8000
RUN chmod 777 /opt
COPY --from=builder /build/notification /usr/bin
USER nobody
ENTRYPOINT ["/usr/bin/notification"]
