#!/usr/bin/env bash

set -ex

cd /tmp

sudo apt-get update -qq && sudo apt-get install -y -qq nmap unzip openjdk-8-jdk-headless

# clojure installation
curl -O https://download.clojure.org/install/linux-install-1.10.0.403.sh
chmod +x linux-install-1.10.0.403.sh
sudo ./linux-install-1.10.0.403.sh
rm linux-install-1.10.0.403.sh
clojure -e '(println "it works!")'

# protoc installation
curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip protoc-3.6.1-linux-x86_64.zip -d protoc3
rm protoc-3.6.1-linux-x86_64.zip
sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/
rm protoc3 -rf
which protoc
