# syntax=docker/dockerfile:1

FROM golang:1.21-bookworm AS build

WORKDIR /app

COPY go.mod ./

COPY  . ./
RUN CGO_ENABLED=0 go build -o turn-conf-generator

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build app/turn-conf-generator /turn-conf-generator

ENV REALM="relrel.org"
ENV TURN_DOMAIN="turn-around.pro"
ENV TURNS_CERT_FILE="/etc/coturn/certs/turn-around.pro.cert"
ENV TURNS_KEY_FILE="/etc/coturn/certs/turn-around.pro.key"
ENV TLS_PORT="5349"

ENTRYPOINT ["/turn-conf-generator"]