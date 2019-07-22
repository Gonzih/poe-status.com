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

RUN ./.circleci/deps.sh
RUN make clean build-binaries release


FROM alpine:latest

RUN apk add --no-cache ca-certificates postgresql-client nmap

WORKDIR /app

COPY --from=builder /go/src/github.com/Gonzih/poe-status.com/cmd/poe-status-client/poe-status-client /bin/
COPY --from=builder /go/src/github.com/Gonzih/poe-status.com/cmd/poe-status-server/poe-status-server /bin/

EXPOSE 8080
