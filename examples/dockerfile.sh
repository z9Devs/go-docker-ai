#!/bin/bash
go build .
#
./go-docker-ai dockerfile create -t golang -p <path>
#
./go-docker-ai dockerfile lint -f <path>
#