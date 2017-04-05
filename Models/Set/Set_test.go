package Set

import (
	"encoding/json"
	"testing"

	//	"gopkg.in/mgo.v2/bson"

	"filip/WeatherStationREST/Models/Brightness"
	"filip/WeatherStationREST/Models/Temperature"
)

var testObjects []Set = GetSetsTestCases()

func TestCreating(t *testing.T) {
	var object *Set
	var objectFloatValue float64
	var expectedFloatValue float64
	var expected []Set = GetSetsTestCases()
	var err error

	for index, c := range getMapSetsTestCases() {
		t.Logf("Creating object '%d'", index)

		object, err = NewFromMap(c)
		if err != nil {
			t.Fatalf("Error when creating: %s", err.Error())
		}

		if *object.Brightness != *expected[index].Brightness {
			t.Fatalf("Expected brightness is '%d' got '%d'", *expected[index].Brightness, *object.Brightness)
		}

		objectFloatValue, _ = object.Temperature.Float64()
		expectedFloatValue, _ = expected[index].Temperature.Float64()
		if objectFloatValue != expectedFloatValue {
			t.Fatalf("Expected temperature is '%v' got '%v'", expectedFloatValue, objectFloatValue)
		}
	}
}

func TestFailNewFromMap(t *testing.T) {
	var err error

	_, err = NewFromMap(map[string]string{"T": "-2"})
	if err.Error() != "No brightness data 'map[T:-2]'" {
		t.Fatalf("Expected error with message\n\t'%s'\n\texpected '%s'", err.Error(), "No brightness data 'map[T:-2]'")
	}

	_, err = NewFromMap(map[string]string{"B": "689"})
	if err.Error() != "No temperature data 'map[B:689]'" {
		t.Fatalf("Expected error with message\n\t'%s'\n\texpected '%s'", err.Error(), "No temperature data 'map[B:689]'")
	}
}

func TestFailParseString(t *testing.T) {
	var err error

	for _, temp := range Temperature.GetUnparsableStringSlice() {
		_, err = ParseStrings(temp, "689")
		if err == nil {
			t.Fatalf("Expected an error for temperature '%s' and brightness '689'", temp)
		}
	}

	for _, brig := range Brightness.GetUnparsableStringSlice() {
		_, err = ParseStrings("-1", brig)
		if err == nil {
			t.Fatalf("Expected an error for temperature '-1' and brightness '%s'", brig)
		}
	}
}

func TestMarshalingUnmarshaling(t *testing.T) {
	var marshal []byte
	var unmarshaled Set
	var unmarshaledFloatValue float64
	var expectedFloatValue float64
	var expected []Set = GetSetsTestCases()
	var err error

	for index, c := range testObjects {
		t.Logf("Marshaling and unmarshaling object '%d'", index)

		marshal, err = json.Marshal(c)
		if err != nil {
			t.Fatalf("Error when marshaling: %s", err.Error())
		}

		err = json.Unmarshal(marshal, &unmarshaled)
		if err != nil {
			t.Fatalf("Error when unmarshaling: %s", err.Error())
		}

		if *unmarshaled.Brightness != *expected[index].Brightness {
			t.Fatalf("Expected brightness is '%d' got '%d'", *expected[index].Brightness, *unmarshaled.Brightness)
		}

		unmarshaledFloatValue, _ = unmarshaled.Temperature.Float64()
		expectedFloatValue, _ = expected[index].Temperature.Float64()
		if unmarshaledFloatValue != expectedFloatValue {
			t.Fatalf("Expected temperature is '%v' got '%v'", expectedFloatValue, unmarshaledFloatValue)
		}
	}
}

func TestFailUnmarshalJSON(t *testing.T) {
	var object Set = Set{}
	var err error = object.UnmarshalJSON([]byte{})

	if err == nil {
		t.Fatalf("Expected error got none")
	}
}

func TestSave(t *testing.T) {
	var err error

	for index, _ := range testObjects {
		t.Logf("Saving object '%d'", index)

		if testObjects[index].IsNew() != true {
			t.Errorf("Object '%d' is not new before save, got 'false', expected 'true'", index)
		}
		if testObjects[index].GetId().String() != "ObjectIdHex(\"\")" {
			t.Errorf("Object '%d' has '%s' ObjectId string, expected 'ObjectIdHex(\"\")'", testObjects[index].GetId().String())
		}

		err = testObjects[index].Save()
		if err != nil {
			t.Fatalf("Error when saving object: %s", err.Error())
		}

		if testObjects[index].IsNew() != false {
			t.Errorf("Object is new after save, got 'true', expected 'false'", index)
		}
		if testObjects[index].GetId().String() == "ObjectIdHex(\"\")" {
			t.Errorf("Object has '%s' ObjectId string, expected different than 'ObjectIdHex(\"\")'", testObjects[index].GetId().String())
		}
	}
}

func TestDelete(t *testing.T) {
	var err error

	for index, _ := range testObjects {
		t.Logf("Deleting object '%d'", index)

		if err != nil {
			t.Fatalf("Error when saving object %s", err.Error())
		}

		if testObjects[index].IsNew() != false {
			t.Errorf("Object is new before deleting, got 'true', expected 'false'", index)
		}
		if testObjects[index].GetId().String() == "ObjectIdHex(\"\")" {
			t.Errorf("Object has '%s' ObjectId string, expected different than 'ObjectIdHex(\"\")'", testObjects[index].GetId().String())
		}

		err = testObjects[index].Delete()
		if err != nil {
			t.Fatalf("Error when deleting object: %s", err.Error())
		}
	}
}

/*func TestFailSetBSON(t *testing.T) {
	var object Set = Set{}
	var err error = object.SetBSON(bson.Raw{})

	if err == nil {
		t.Fatalf("Expected error got none")
	}
}*/
