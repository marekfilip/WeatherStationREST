package Brightness

import (
	"encoding/json"
	"math/big"
	"time"

	conf "filip/WeatherStationREST/Config"
	db "filip/WeatherStationREST/DbConnection/Mongo"

	"gopkg.in/mgo.v2/bson"
)

type Brightness struct {
	Id       bson.ObjectId
	Created  time.Time
	Modified time.Time

	Timestamp time.Time
	Value     *big.Float

	exists bool
}

func New(stringValue string) (*Brightness, error) {
	floatValue, _, err := big.ParseFloat(stringValue, 10, 0, big.ToNearestEven)

	if err != nil {
		return nil, err
	}

	return &Brightness{
		Timestamp: time.Now(),
		Value:     floatValue,
	}, nil
}

func (t *Brightness) Save() error {
	connection, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = connection.Collection(conf.GetBrightnessCollectionName()).Save(t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Brightness) Delete() error {
	connection, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = connection.Collection(conf.GetBrightnessCollectionName()).DeleteDocument(t)
	if err != nil {
		return err
	}

	return nil
}

// Implements json.Marshaler
func (t *Brightness) MarshalJSON() ([]byte, error) {
	value, _ := t.Value.Float64()

	return json.Marshal(&struct {
		Timestamp int64   `json:"timestamp"`
		Value     float64 `json:"value"`
	}{
		Timestamp: t.Timestamp.Unix(),
		Value:     value,
	})
}

// Implements json.Unmarshaler
func (t *Brightness) UnmarshalJSON(data []byte) error {
	var temporaryObject struct {
		Timestamp int64   `json:"timestamp"`
		Value     float64 `json:"value"`
	}

	err := json.Unmarshal(data, &temporaryObject)
	if err != nil {
		return err
	}

	*t = Brightness{
		Timestamp: time.Unix(temporaryObject.Timestamp, 0),
		Value:     big.NewFloat(temporaryObject.Value),
	}

	return nil
}

// Satisfy the bson.Getter
func (t Brightness) GetBSON() (interface{}, error) {
	value, _ := t.Value.Float64()

	return struct {
		Id        bson.ObjectId `bson:"_id,omitempty"`
		Created   int64         `bson:"_created"`
		Modified  int64         `bson:"_modified"`
		Timestamp int64         `bson:"timestamp"`
		Value     float64       `bson:"value"`
	}{
		Id:        t.GetId(),
		Created:   t.Created.Unix(),
		Modified:  t.Modified.Unix(),
		Timestamp: t.Timestamp.Unix(),
		Value:     value,
	}, nil
}

// Satisfy the bson.Setter
func (t *Brightness) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id        bson.ObjectId `bson:"_id,omitempty"`
		Created   int64         `bson:"_created"`
		Modified  int64         `bson:"_modified"`
		Timestamp int64         `bson:"timestamp"`
		Value     float64       `bson:"value"`
	})

	bsonErr := raw.Unmarshal(decoded)

	if bsonErr == nil {
		*t = Brightness{
			Id:        decoded.Id,
			Created:   time.Unix(decoded.Created, 0),
			Modified:  time.Unix(decoded.Modified, 0),
			Timestamp: time.Unix(decoded.Timestamp, 0),
			Value:     big.NewFloat(decoded.Value),
		}
		return nil
	} else {
		return bsonErr
	}
}

// Satisfy the bongo.NewTracker
func (t *Brightness) SetIsNew(isNew bool) {
	t.exists = !isNew
}

func (t *Brightness) IsNew() bool {
	return !t.exists
}

// Satisfy the bongo.Document
func (t *Brightness) GetId() bson.ObjectId {
	return t.Id
}

func (t *Brightness) SetId(id bson.ObjectId) {
	t.Id = id
}

// Satisfy the bongo.TimeTracker
func (t *Brightness) SetCreated(ts time.Time) {
	t.Created = ts
}

func (t *Brightness) SetModified(ts time.Time) {
	t.Modified = ts
}
