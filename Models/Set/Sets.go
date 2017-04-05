package Set

import (
	"sort"
	"time"

	"gopkg.in/mgo.v2/bson"

	db "filip/WeatherStationREST/DbConnection/Mongo"
)

type Sets struct {
	List      *[]Set
	StartTime time.Time
	EndTime   time.Time
}

func NewSets() *Sets {
	return &Sets{
		List:      new([]Set),
		StartTime: time.Time{},
		EndTime:   time.Time{},
	}
}

func (sets *Sets) Append(object Set) {
	var zeroTime = time.Time{}
	*sets.List = append(*sets.List, object)

	if object.Created.Before(sets.StartTime) || sets.StartTime == zeroTime {
		sets.StartTime = object.Created
	}

	if object.Created.After(sets.EndTime) || sets.EndTime == zeroTime {
		sets.EndTime = object.Created
	}
}

func Find(startTime, endTime time.Time) (*Sets, error) {
	results := db.GetSetCollection().Find(
		bson.M{
			"_created": bson.M{
				"$gte": startTime.Unix(),
				"$lte": endTime.Unix(),
			}})

	var temp Set
	var list *Sets = NewSets()
	for results.Next(&temp) {
		list.Append(temp)
	}

	sort.Sort(list)

	return list, nil
}

func (list Sets) Len() int {
	return len(*list.List)
}

func (list Sets) Less(i, j int) bool {
	return (*list.List)[i].Created.Before((*list.List)[j].Created)
}

func (list Sets) Swap(i, j int) {
	(*list.List)[i], (*list.List)[j] = (*list.List)[j], (*list.List)[i]
}
