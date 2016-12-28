package main

import (
	"log"
	//"math/big"
	"time"

	temp "filip/WeatherStationREST/Models/Temperature"
)

func main() {
	/*var obj temp.Temperature = temp.Temperature{
		Timestamp: time.Unix(time.Now().Unix(), 0),
		Value:     big.NewFloat(13.123),
	}

	var i int64
	for i = 1; i <= 100; i++ {
		obj = temp.Temperature{
			Timestamp: time.Unix(obj.Timestamp.Unix()+i, 0),
			Value:     big.NewFloat(13.123),
		}

		t0 := time.Now()
		err := obj.Save()
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Printf("Object %d saved in %d ms\n", i, (time.Now().Sub(t0).Nanoseconds() / 1e6))
		}
	}*/

	obj := new(temp.Temperatures)

	obj.Find(time.Date(2016, time.December, 23, 8, 0, 0, 0, time.Now().Location()), time.Now())

	for _, one := range *obj {
		log.Printf("%T %+v\n", one.Timestamp, one.Timestamp)
		log.Printf("%T %+v\n", one.Value, one.Value)
	}
}
