package Signed

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"

	db "filip/WeatherStationREST/DbConnection/Mongo"
	"filip/WeatherStationREST/Models/Set"
	"filip/WeatherStationREST/Models/Temperature"
)

type SignedTemperature struct {
	Temperature *Temperature.Temperature
	Timestamp   time.Time
}

// Implements json.Marshaler
func (s SignedTemperature) MarshalJSON() ([]byte, error) {
	temperature, _ := s.Temperature.Float64()
	return json.Marshal(&struct {
		Timestamp   int64   `json:"timestamp"`
		Temperature float64 `json:"temperature"`
	}{
		Timestamp:   s.Timestamp.Unix(),
		Temperature: temperature,
	})
}

func (s *SignedTemperature) UnmarshalJSON(data []byte) error {
	var decoded struct {
		Timestamp   int64   `json:"timestamp"`
		Temperature float64 `json:"temperature"`
	}

	err := json.Unmarshal(data, &decoded)
	if err != nil {
		return err
	}

	*s = SignedTemperature{
		Timestamp:   time.Unix(decoded.Timestamp, 0),
		Temperature: Temperature.NewTemperature(decoded.Temperature),
	}

	return nil
}

func GetLast() (*SignedTemperature, error) {
	var tempSet Set.Set

	results := db.GetSetCollection().Find(bson.M{
		"_created": bson.M{
			"$gte": time.Now().Add(time.Duration(-10) * time.Minute).Unix(),
		}})
	results.Query.Sort("-$natural").Limit(1)

	hasNext := results.Next(&tempSet)

	if !hasNext && results.Error != nil {
		return nil, results.Error
	}

	return &SignedTemperature{
		Timestamp:   tempSet.Created,
		Temperature: tempSet.Temperature,
	}, nil
}
