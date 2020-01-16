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
      sudo

WORKDIR /go/src/github.com/Gonzih/poe-status.com

ENV GO111MODULE=on

COPY . .

RUN ./.github/deps.sh
RUN make clean release-binaries


FROM ubuntu:latest

# RUN apk add --no-cache ca-certificates postgresql-client nmap
RUN apt-get update -qq && apt-get install -y -qq  ca-certificates nmap iputils-ping

COPY --from=builder /go/src/github.com/Gonzih/poe-status.com/cmd/poe-status-client/poe-status-client /usr/bin/
COPY --from=builder /go/src/github.com/Gonzih/poe-status.com/cmd/poe-status-server/poe-status-server /usr/bin/

EXPOSE 8080
