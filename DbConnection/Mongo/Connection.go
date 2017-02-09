package Mongo

import (
	conf "filip/WeatherStationREST/Config"

	"github.com/maxwellhealth/bongo"
)

var instance *bongo.Connection = nil

func GetConnection() *bongo.Connection {
	if instance == nil {
		newInstance, _ := bongo.Connect(conf.GetBongoConfig())
		instance = newInstance
	}

	return instance
}
