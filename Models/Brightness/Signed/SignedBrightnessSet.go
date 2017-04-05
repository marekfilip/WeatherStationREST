package Signed

import (
	"sort"
	"time"

	"gopkg.in/mgo.v2/bson"

	db "filip/WeatherStationREST/DbConnection/Mongo"
	"filip/WeatherStationREST/Models/Set"
)

// BrightnessSet is a slice of signed Brightness
type BrightnessSet []Brightness

// NewBrightnessSet returns an empty BrightnessSet
func NewBrightnessSet() BrightnessSet {
	return make(BrightnessSet, 0)
}

// Append adds brightness object to BrightnessSet
func (list *BrightnessSet) Append(object Brightness) {
	*list = append(*list, object)
}

// Deletes element with specified index
func (list *BrightnessSet) Delete(index uint) {
	*list = append((*list)[:index], (*list)[index+1:]...)
}

// Find gets brightness set from DB for specified start and end time
func Find(startTime, endTime time.Time) (BrightnessSet, error) {
	results := db.GetSetCollection().Find(
		bson.M{
			"_created": bson.M{
				"$gte": startTime.Unix(),
				"$lte": endTime.Unix(),
			}})

	var tempSet Set.Set
	var list = NewBrightnessSet()
	for results.Next(&tempSet) {
		list.Append(Brightness{
			Brightness: tempSet.Brightness,
			Timestamp:  tempSet.Created,
		})
	}

	sort.Sort(list)

	return list, nil
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
