package main

import (
	"log"
	"os"
	"time"

	"filip/WeatherStationREST/Models/Set"
	"filip/WeatherStationREST/REST"
	"filip/WeatherStationREST/SerialRead"
)

var environment string = os.Getenv("WEATHER_STATION_ENV")

func main() {
	checkIsEnvironmentSet()

	var device *SerialRead.SerialRead = SerialRead.Init()
	var rest *REST.WeatherStationREST
	var err error

	go func() {
		for {
			data := device.GetData()
			if err != nil {
				log.Fatal(err.Error())
			}

			composition, err := Set.NewFromMap(data)
			if err != nil {
				log.Fatal(err.Error())
				return
			}

			composition.Save()
			<-time.After(time.Duration(2) * time.Minute)
		}
	}()

	rest, err = REST.New(REST.DevStack)
	if err != nil {
		log.Fatal(err.Error())
	}

	rest.Start()
}

func checkIsEnvironmentSet() {
	if environment == "" {
		log.Println("Set environment in WEATHER_STATION_ENV")
		os.Exit(1)
	}
}
