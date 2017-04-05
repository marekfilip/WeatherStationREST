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
	log.SetOutput(os.Stdout)
	checkIsEnvironmentSet()

	var device *SerialRead.SerialRead = SerialRead.Init()
	var rest *REST.WeatherStationREST
	var err error

	go func() {
		for {
			<-time.After(time.Duration(15) * time.Second)

			data, err := device.GetData()
			if err != nil {
				log.Println(err.Error())
				continue
			}

			composition, err := Set.NewFromMap(data)
			if err == nil {
				err = composition.Save()

				if err != nil {
					log.Println(err.Error())
				}
			} else {
				log.Printf(
					"Error when creating set\n\tError: %s\n\tData: %v\n",
					err.Error(),
					data,
				)
			}

			<-time.After(time.Duration(2) * time.Minute)
		}
	}()

	rest, err = REST.New(REST.DevStack)
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatalln(err.Error())
	}

	rest.Start()
}

func checkIsEnvironmentSet() {
	if environment == "" {
		log.Fatalln("Set environment in WEATHER_STATION_ENV")
	}
}
