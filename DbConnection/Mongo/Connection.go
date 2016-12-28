package Mongo

import (
	conf "filip/WeatherStationREST/Config"
	"github.com/maxwellhealth/bongo"
)

var instance *bongo.Connection = nil

func GetConnection() (*bongo.Connection, error) {
	if instance == nil {
		newInstance, err := bongo.Connect(conf.GetBongoConfig())
		if err != nil {
			return nil, err
		}
		instance = newInstance
	}

	return instance, nil
}
