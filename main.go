package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"filip/WeatherStationREST/Models/Set"
	"filip/WeatherStationREST/REST"
	sr "filip/WeatherStationREST/SerialRead"
	"filip/tracer"
)

var environment string = os.Getenv("WEATHER_STATION_ENV")

func main() {
	checkIsEnvironmentSet()

	var device *sr.SerialRead = sr.Init()
	var rest *REST.WeatherStationREST
	var err error

	// Tracer dla odczytów
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
	}
	folderPath := strings.Join([]string{wd, "logs"}, "/")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0755)
	}
	file, err := os.OpenFile(strings.Join([]string{folderPath, "serial_read.log"}, "/"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()
	device.SetTracer(tracer.New(file))

	go func() {
		for {
			<-time.After(time.Duration(15) * time.Second)

			data, err := GetData(device)
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
		log.Fatalln(err.Error())
	}

	rest.Start()
}

func GetData(sr *sr.SerialRead) (map[string]string, error) {
	var output map[string]string = map[string]string{}

	tries := 0
	for {
		var data string = string(sr.ReadData())
		var splitedData = strings.Split(data, ";")

		for _, value := range splitedData {
			tmp := strings.Split(strings.TrimSpace(value), ":")

			if len(tmp) == 2 {
				output[tmp[0]] = tmp[1]
			}
		}

		if output["T"] != "-273.15" && output["B"] != "" {
			break
		}

		tries++
		if tries >= 3 {
			return nil, fmt.Errorf("Tries count too hi")
		}
	}

	return output, nil
}

func checkIsEnvironmentSet() {
	if environment == "" {
		log.Fatalln("Set environment in WEATHER_STATION_ENV")
	}
}
