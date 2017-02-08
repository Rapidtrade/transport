#!/bin/bash

echo "============== Cleaning up before starting =============="
go clean
rm -rf $GOPATH/bin/transport

echo "============== Building binary to be used in the container =============="
env GOOS=linux GOARCH=arm go build -v github.com/Rapidtrade/transport
cp $GOPATH/bin/transport $GOPATH/src/github.com/Rapidtrade/transport

echo "============== Starting docker build =============="
docker build . -t transport

echo "============== Cleaning Up =============="
go clean
