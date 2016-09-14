#!/usr/bin/env bash
cd src
echo "Building GO application..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../deployment/ipLookup main.go
#GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o deployment/ipLookup.exe src/gonp/main.go
echo "Done building loveliness"