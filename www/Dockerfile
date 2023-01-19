FROM alpine:latest

RUN apk add --no-cache --update npm=8.5.1

RUN mkdir -p /.npm && chmod -R 777 /.npm
ADD . /src
RUN chmod -R 777 /src

WORKDIR /src
RUN npm install

EXPOSE 8080
ENTRYPOINT exec npm run serve
