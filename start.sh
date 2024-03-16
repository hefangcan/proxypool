#!/bin/sh

# download
wget https://github.com/jth445600/hello-world/raw/master/proxypool-linux-amd64/proxypoolv0.7.3 -O proxypool
chmod 755 proxypool
wget https://github.com/jth445600/hello-world/raw/master/proxypool-linux-amd64/source.yaml -O source.yaml
wget https://github.com/jth445600/hello-world/raw/master/proxypool-linux-amd64/config.yaml -O config.yaml
wget https://github.com/jth445600/hello-world/raw/master/proxypool-linux-amd64/assets.zip -O assets.zip
unzip -d assets assets.zip

# storefiles

# start


./proxypool -c config.yaml &

