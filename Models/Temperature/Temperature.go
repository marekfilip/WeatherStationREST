package Temperature

import (
	"encoding/json"
	"math/big"
	"time"

	conf "filip/WeatherStationREST/Config"
	db "filip/WeatherStationREST/DbConnection/Mongo"

	"gopkg.in/mgo.v2/bson"
)

type Temperature struct {
	Id        bson.ObjectId
	Created   time.Time
	Modified  time.Time
	Timestamp time.Time
	Value     *big.Float
	exists    bool
}

func (t *Temperature) Save() error {
	connection, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = connection.Collection(conf.GetTemperatureCollectionName()).Save(t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Temperature) Delete() error {
	connection, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = connection.Collection(conf.GetTemperatureCollectionName()).DeleteDocument(t)
	if err != nil {
		return err
	}

	return nil
}

// Implements json.Marshaler
func (t *Temperature) MarshalJSON() ([]byte, error) {
	value, _ := t.Value.Float64()

	return json.Marshal(&struct {
		Id        bson.ObjectId `json:"_id"`
		Created   int64         `json:"_created"`
		Modified  int64         `json:"_modified"`
		Timestamp int64         `json:"timestamp"`
		Value     float64       `json:"value"`
	}{
		Id:        t.GetId(),
		Created:   t.Created.Unix(),
		Modified:  t.Modified.Unix(),
		Timestamp: t.Timestamp.Unix(),
		Value:     value,
	})
}

// Implements json.Unmarshaler
func (t *Temperature) UnmarshalJSON(data []byte) error {
	var temporaryObject struct {
		Id        bson.ObjectId `json:"_id"`
		Created   int64         `json:"_created"`
		Modified  int64         `json:"_modified"`
		Timestamp int64         `json:"timestamp"`
		Value     float64       `json:"value"`
	}

	err := json.Unmarshal(data, &temporaryObject)
	if err != nil {
		return err
	}

	*t = Temperature{
		Id:        temporaryObject.Id,
		Created:   time.Unix(temporaryObject.Created, 0),
		Modified:  time.Unix(temporaryObject.Modified, 0),
		Timestamp: time.Unix(temporaryObject.Timestamp, 0),
		Value:     big.NewFloat(temporaryObject.Value),
	}

	return nil
}

// Satisfy the bson.Getter
func (t Temperature) GetBSON() (interface{}, error) {
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
func (t *Temperature) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id        bson.ObjectId `bson:"_id,omitempty"`
		Created   int64         `bson:"_created"`
		Modified  int64         `bson:"_modified"`
		Timestamp int64         `bson:"timestamp"`
		Value     float64       `bson:"value"`
	})

	bsonErr := raw.Unmarshal(decoded)

	if bsonErr == nil {
		*t = Temperature{
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
func (t *Temperature) SetIsNew(isNew bool) {
	t.exists = !isNew
}

func (t *Temperature) IsNew() bool {
	return !t.exists
}

// Satisfy the bongo.Document
func (t *Temperature) GetId() bson.ObjectId {
	return t.Id
}

func (t *Temperature) SetId(id bson.ObjectId) {
	t.Id = id
}

// Satisfy the bongo.TimeTracker
func (t *Temperature) SetCreated(ts time.Time) {
	t.Created = ts
}

func (t *Temperature) SetModified(ts time.Time) {
	t.Modified = ts
}
