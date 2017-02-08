package Brightness

import (
	"time"

	conf "filip/WeatherStationREST/Config"
	db "filip/WeatherStationREST/DbConnection/Mongo"
	"gopkg.in/mgo.v2/bson"
)

type BrightnessSet []Brightness

func (list *BrightnessSet) Append(object Brightness) {
	*list = append(*list, object)
}

func (list *BrightnessSet) Find(startTime, stopTime time.Time) error {
	connection, err := db.GetConnection()
	if err != nil {
		return err
	}

	results := connection.Collection(conf.GetBrightnessCollectionName()).Find(
		bson.M{
			"timestamp": bson.M{
				"$gte": startTime.Unix(),
				"$lte": stopTime.Unix(),
			}})

	var temp Brightness
	for results.Next(&temp) {
		list.Append(temp)
	}

	return nil
}

func (list BrightnessSet) Len() int {
	return len(list)
}

func (list BrightnessSet) Less(i, j int) bool {
	return list[i].Timestamp.Before(list[j].Timestamp)
}

func (list BrightnessSet) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
