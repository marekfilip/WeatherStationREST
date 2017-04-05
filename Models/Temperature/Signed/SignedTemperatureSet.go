package Signed

import (
	"sort"
	"time"

	"gopkg.in/mgo.v2/bson"

	db "filip/WeatherStationREST/DbConnection/Mongo"
	"filip/WeatherStationREST/Models/Set"
)

type SignedTemperatureSet []SignedTemperature

func NewSignedTemperatureSet() SignedTemperatureSet {
	return make(SignedTemperatureSet, 0)
}

func (list *SignedTemperatureSet) Append(object SignedTemperature) {
	*list = append(*list, object)
}

func Find(startTime, endTime time.Time) (SignedTemperatureSet, error) {
	results := db.GetSetCollection().Find(
		bson.M{
			"_created": bson.M{
				"$gte": startTime.Unix(),
				"$lte": endTime.Unix(),
			}})

	var tempSet Set.Set
	var list SignedTemperatureSet = NewSignedTemperatureSet()
	for results.Next(&tempSet) {
		list.Append(SignedTemperature{
			Temperature: tempSet.Temperature,
			Timestamp:   tempSet.Created,
		})
	}

	sort.Sort(list)

	return list, nil
}

func (list SignedTemperatureSet) Len() int {
	return len(list)
}

func (list SignedTemperatureSet) Less(i, j int) bool {
	return list[i].Timestamp.Before(list[j].Timestamp)
}

func (list SignedTemperatureSet) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
