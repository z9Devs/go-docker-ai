#!/bin/bash
go build .
./Go-DockerLint-Ai dockerlint -f Dockerfile          
