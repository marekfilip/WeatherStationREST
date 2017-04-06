#!/bin/bash

export WEATHER_STATION_ENV="test"
go test -cover $(go list ./...)