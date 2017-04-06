#!/bin/bash

export WEATHER_STATION_ENV="test"
for i in $(go list ./...); do
	go test -cover -coverprofile=coverage.out $i
	if [[ -f coverage.out ]]; then
		go tool cover -func=coverage.out && go tool cover -html=coverage.out && rm coverage.out
	fi
done
