#! /bin/bash

rm -rf static
webpack
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o ignite-admin
docker build -t $1 .
