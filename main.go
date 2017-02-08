package main

import (
	"log"
	"time"

	"filip/WeatherStationREST/Models/Composition"
	"filip/WeatherStationREST/REST"
	"filip/WeatherStationREST/SerialRead"
)

func main() {
	var device *SerialRead.SerialRead = SerialRead.Init()
	var rest *REST.WeatherStationREST
	var err error

	go func() {
		for {
			data := device.GetData()
			if err != nil {
				log.Fatal(err.Error())
			}

			composition, err := Composition.GetCompositionFromSerialData(data)
			if err != nil {
				log.Fatal(err.Error())
				return
			}

			composition.SaveAll()
			<-time.After(time.Duration(2) * time.Minute)
		}
	}()

	rest, err = REST.New(REST.DevStack)
	if err != nil {
		log.Fatal(err.Error())
	}

	rest.Start()
}
