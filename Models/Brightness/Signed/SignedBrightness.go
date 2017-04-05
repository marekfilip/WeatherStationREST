package Signed

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"

	db "filip/WeatherStationREST/DbConnection/Mongo"
	bri "filip/WeatherStationREST/Models/Brightness"
	"filip/WeatherStationREST/Models/Set"
)

// Brightness is a Brightness signed with timestamp
type Brightness struct {
	Brightness *bri.Brightness
	Timestamp  time.Time
}

// MarshalJSON implements json.Marshaler
func (s Brightness) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Timestamp  int64  `json:"timestamp"`
		Brightness uint16 `json:"brightness"`
	}{
		Timestamp:  s.Timestamp.Unix(),
		Brightness: uint16(*s.Brightness),
	})
}

// UnmarshalJSON implements json.Unmarshaler
func (s *Brightness) UnmarshalJSON(data []byte) error {
	var decoded struct {
		Timestamp  int64  `json:"timestamp"`
		Brightness uint16 `json:"brightness"`
	}

	err := json.Unmarshal(data, &decoded)
	if err != nil {
		return err
	}

	*s = Brightness{
		Timestamp:  time.Unix(decoded.Timestamp, 0),
		Brightness: bri.NewBrightness(decoded.Brightness),
	}

	return nil
}

// GetLast gets the last brightness reading
func GetLast() (*Brightness, error) {
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

	return &Brightness{
		Timestamp:  tempSet.Created,
		Brightness: tempSet.Brightness,
	}, nil
}
