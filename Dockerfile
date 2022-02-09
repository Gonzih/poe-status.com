FROM golang:latest AS builder


RUN apt-get update -qq && \
    apt-get -y -qq install \
      bash \
      ca-certificates \
      curl \
      gcc \
      git \
      make \
      postgresql-client \
      unzip \
      openjdk-11-jre \
      sudo

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN make deps install-clojure
RUN make clean release-binaries


FROM ubuntu:latest

# RUN apk add --no-cache ca-certificates postgresql-client nmap
RUN apt-get update -qq && apt-get install -y -qq  ca-certificates nmap iputils-ping

COPY --from=builder /app/cmd/poe-status-client/poe-status-client /usr/bin/
COPY --from=builder /app/cmd/poe-status-server/poe-status-server /usr/bin/

EXPOSE 8080
