#!/bin/sh
#
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o go-docker-ai-linux-amd64 .
#CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -a -o go-docker-ai-linux-arm64 .
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -a -o go-docker-ai-windows-amd64 .
#CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -a -o go-docker-ai-windows-arm64 .
#
mkdir -p /tmp/bin
pwd
ls -al
tar -czvf /tmp/bin/go-docker-ai.tar.gz --transform='s|.*/||' ./go-docker-ai-*
#
ls -l /tmp/bin