#!/usr/bin/env bash
cd src
go get ./...
cd ../deployment
go run ../src/main.go