#!/bin/bash

if [[ ! -d "./bin" ]]; then
	mkdir ./bin
fi

env GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ./bin/weatherread main.go