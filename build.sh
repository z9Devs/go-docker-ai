#!/bin/sh
#
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o go-dockerlint-ai-linux-amd64 .
#CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -a -o go-dockerlint-ai-linux-arm64 .
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -a -o go-dockerlint-ai-windows-amd64 .
#CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -a -o go-dockerlint-ai-windows-arm64 .
#
mkdir -p /tmp/bin
pwd
ls -al
tar -czvf /tmp/bin/go-dockerlint-ai.tar.gz --transform='s|.*/||' ./go-dockerlint-ai-*
#
ls -l /tmp/bin