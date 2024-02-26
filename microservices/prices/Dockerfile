# Builder
FROM golang:1.19-alpine AS builder
ADD . /src/prices
WORKDIR /src/prices
RUN [ -x "$(command -v oapi-codegen)" ] || go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
    mkdir -p autogen && \
    sed -e 's/^/    /' README.md > autogen/README.md && \
    sed -e '/    %README.md%/{' -e "r autogen/README.md" -e 'd' -e '}' oapi3.yaml > autogen/oapi3.yaml && \
    oapi-codegen -o autogen/server.go -old-config-style -package autogen -generate chi-server autogen/oapi3.yaml && \
    oapi-codegen -o autogen/types.go  -old-config-style -package autogen -generate types      autogen/oapi3.yaml && \
    oapi-codegen -o autogen/client.go -old-config-style -package autogen -generate client     autogen/oapi3.yaml
RUN CGO_ENABLED=1 GOOS=linux go build -o /build/prices .

# Service
FROM alpine:latest
EXPOSE 8000
RUN chmod 777 /opt
COPY --from=builder /build/prices /usr/bin
USER nobody
ENTRYPOINT ["/usr/bin/prices"]
