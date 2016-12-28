package Temperature

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"
)

var temperatureTestCases = getTestCases()

func TestCreateTemperature(t *testing.T) {
	var value float64

	for index, c := range temperatureTestCases {
		temperatureTestCases[index].object = &Temperature{
			Timestamp: time.Unix(c.timestamp, 0),
			Value:     big.NewFloat(c.temperature),
		}
	}

	for index, c := range temperatureTestCases {
		t.Logf("Creating object '%d'", index)

		if c.object.Timestamp.Unix() != c.timestamp {
			t.Fatalf("Expected timestamp is '%d' got '%d'", index, c.timestamp, c.object.Timestamp.Unix())
		}

		value, _ = c.object.Value.Float64()
		if value != c.temperature {
			t.Fatalf("Expected value is '%f' got '%f'", index, c.temperature, value)
		}
	}
}

func TestTemperatureMarshaling(t *testing.T) {
	var marshal []byte
	var unmarshaled Temperature
	var err error

	for index, c := range temperatureTestCases {
		t.Logf("Marshaling and unmarshaling object '%d'", index)

		marshal, err = json.Marshal(&c.object)
		if err != nil {
			t.Fatalf("Error when marshaling: %s", err.Error())
		}

		err = json.Unmarshal(marshal, &unmarshaled)
		if err != nil {
			t.Fatalf("Error when unmarshaling: %s", err.Error())
		}

		if unmarshaled.Timestamp != c.object.Timestamp {
			t.Fatalf("Expected timestamp is '%d' got '%d'", index, unmarshaled.Timestamp, c.object.Timestamp)
		}

		if unmarshaled.Value.Cmp(c.object.Value) != 0 {
			t.Fatalf("Expected value is '%+v' got '%+v'", index, *unmarshaled.Value, *c.object.Value)
		}
	}
}

func TestSaveTemperature(t *testing.T) {
	var err error

	for index, c := range temperatureTestCases {
		t.Logf("Saving object '%d'", index)

		if c.object.IsNew() != true {
			t.Errorf("Object '%d' is not new before save, got 'false', expected 'true'", index)
		}
		if c.object.GetId().String() != "ObjectIdHex(\"\")" {
			t.Errorf("Object '%d' has '%s' ObjectId string, expected 'ObjectIdHex(\"\")'", c.object.GetId().String())
		}

		err = c.object.Save()
		if err != nil {
			t.Fatalf("Error when saving object: %s", err.Error())
		}

		if c.object.IsNew() != false {
			t.Errorf("Object is new after save, got 'true', expected 'false'", index)
		}
		if c.object.GetId().String() == "ObjectIdHex(\"\")" {
			t.Errorf("Object has '%s' ObjectId string, expected different than 'ObjectIdHex(\"\")'", c.object.GetId().String())
		}
	}
}

func TestDeleteTemperature(t *testing.T) {
	var err error

	for index, c := range temperatureTestCases {
		t.Logf("Deleting object '%d'", index)

		if err != nil {
			t.Fatalf("Error when saving object %s", err.Error())
		}

		if c.object.IsNew() != false {
			t.Errorf("Object is new before deleting, got 'true', expected 'false'", index)
		}
		if c.object.GetId().String() == "ObjectIdHex(\"\")" {
			t.Errorf("Object has '%s' ObjectId string, expected different than 'ObjectIdHex(\"\")'", c.object.GetId().String())
		}

		err = c.object.Delete()
		if err != nil {
			t.Fatalf("Error when deleting object: %s", err.Error())
		}
	}
}
