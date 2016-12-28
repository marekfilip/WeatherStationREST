package Temperature

import (
	"time"

	conf "filip/WeatherStationREST/Config"
	db "filip/WeatherStationREST/DbConnection/Mongo"
	"gopkg.in/mgo.v2/bson"
)

type Temperatures []Temperature

func (list *Temperatures) Append(object Temperature) {
	*list = append(*list, object)
}

func (list *Temperatures) Find(startTime, stopTime time.Time) error {
	connection, err := db.GetConnection()
	if err != nil {
		return err
	}

	results := connection.Collection(conf.GetTemperatureCollectionName()).Find(
		bson.M{
			"timestamp": bson.M{
				"$gte": startTime.Unix(),
				"$lte": stopTime.Unix(),
			}})

	var temp Temperature
	for results.Next(&temp) {
		list.Append(temp)
	}

	return nil
}
