package Set

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	conf "filip/WeatherStationREST/Config"
	db "filip/WeatherStationREST/DbConnection/Mongo"
	"gopkg.in/mgo.v2/bson"

	"filip/WeatherStationREST/Models/Brightness"
	"filip/WeatherStationREST/Models/Temperature"
)

type Set struct {
	Id       bson.ObjectId
	Created  time.Time
	Modified time.Time
	exists   bool

	Brightness  *Brightness.Brightness   `json:"brightness"`
	Temperature *Temperature.Temperature `json:"temperature"`
}

func NewFromMap(data map[string]string) (*Set, error) {
	if _, ok := data["T"]; !ok {
		return nil, fmt.Errorf("No temperature data '%+v'", data)
	}

	if _, ok := data["B"]; !ok {
		return nil, fmt.Errorf("No brightness data '%+v'", data)
	}

	return ParseStrings(data["T"], data["B"])
}

func ParseStrings(t, b string) (*Set, error) {
	newTemperature, err := Temperature.ParseString(t)
	if err != nil {
		return nil, err
	}

	newBrightness, err := Brightness.ParseString(b)
	if err != nil {
		return nil, err
	}

	return New(newTemperature, newBrightness), nil
}

func New(t *Temperature.Temperature, b *Brightness.Brightness) *Set {
	return &Set{
		Temperature: t,
		Brightness:  b,
	}
}

func (s *Set) Save() error {
	var err error = db.GetConnection().Collection(conf.GetSetCollectionName()).Save(s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Set) Delete() error {
	var err error = db.GetConnection().Collection(conf.GetSetCollectionName()).DeleteDocument(s)
	if err != nil {
		return err
	}

	return nil
}

// Implements json.Marshaler
func (s Set) MarshalJSON() ([]byte, error) {
	temperature, _ := s.Temperature.Float64()
	return json.Marshal(&struct {
		Timestamp   int64   `json:"timestamp"`
		Brightness  uint16  `json:"brightness"`
		Temperature float64 `json:"temperature"`
	}{
		Timestamp:   s.Created.Unix(),
		Brightness:  uint16(*s.Brightness),
		Temperature: temperature,
	})
}

func (s *Set) UnmarshalJSON(data []byte) error {
	var decoded struct {
		Timestamp   int64   `json:"timestamp"`
		Brightness  uint16  `json:"brightness"`
		Temperature float64 `json:"temperature"`
	}

	err := json.Unmarshal(data, &decoded)
	if err != nil {
		return err
	}

	*s = Set{
		Created:     time.Unix(decoded.Timestamp, 0),
		Brightness:  Brightness.NewBrightness(decoded.Brightness),
		Temperature: Temperature.NewTemperature(decoded.Temperature),
	}

	return nil
}

// Satisfy the bson.Getter
func (s Set) GetBSON() (interface{}, error) {
	temperature, _ := s.Temperature.Float64()

	return struct {
		Id          bson.ObjectId `bson:"_id,omitempty"`
		Created     int64         `bson:"_created"`
		Modified    int64         `bson:"_modified"`
		Brightness  uint16        `json:"brightness"`
		Temperature float64       `json:"temperature"`
	}{
		Id:          s.GetId(),
		Created:     s.Created.Unix(),
		Modified:    s.Modified.Unix(),
		Brightness:  uint16(*s.Brightness),
		Temperature: temperature,
	}, nil
}

// Satisfy the bson.Setter
func (s *Set) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id          bson.ObjectId `bson:"_id,omitempty"`
		Created     int64         `bson:"_created"`
		Modified    int64         `bson:"_modified"`
		Brightness  uint16        `json:"brightness"`
		Temperature float64       `json:"temperature"`
	})

	bsonErr := raw.Unmarshal(decoded)
	if bsonErr == nil {
		brightness := Brightness.Brightness(decoded.Brightness)
		temperature := Temperature.Temperature(*big.NewFloat(decoded.Temperature))

		*s = Set{
			Id:          decoded.Id,
			Created:     time.Unix(decoded.Created, 0),
			Modified:    time.Unix(decoded.Modified, 0),
			Brightness:  &brightness,
			Temperature: &temperature,
		}
		return nil
	}

	return bsonErr
}

// Satisfy the bongo.NewTracker
func (s *Set) SetIsNew(isNew bool) {
	s.exists = !isNew
}

func (s *Set) IsNew() bool {
	return !s.exists
}

// Satisfy the bongo.Document
func (s *Set) GetId() bson.ObjectId {
	return s.Id
}

func (s *Set) SetId(id bson.ObjectId) {
	s.Id = id
}

// Satisfy the bongo.TimeTracker
func (s *Set) SetCreated(ts time.Time) {
	s.Created = ts
}

func (s *Set) SetModified(ts time.Time) {
	s.Modified = ts
}
