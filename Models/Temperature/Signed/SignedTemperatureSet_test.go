package Signed

import (
	"testing"
	"time"

	"filip/WeatherStationREST/Models/Set"
	"filip/WeatherStationREST/Models/Temperature"
)

func TestFind(t *testing.T) {
	var testSetsObjects []Set.Set = Set.GetSetsTestCases()

	for index, _ := range testSetsObjects {
		testSetsObjects[index].Save()
	}

	signedTemperatures, err := Find(testSetsObjects[0].Created, testSetsObjects[len(testSetsObjects)-1].Created)
	if err != nil {
		t.Fatalf("Error occured:\n%s", err.Error())
	}

	for _, signedTemperature := range signedTemperatures {
	Loop:
		for _, testSetObject := range testSetsObjects {
			if signedTemperature.Timestamp.Unix() == testSetObject.Created.Unix() {
				break Loop
			}

			t.Fatalf("Could not find object with timestamp '%v'", signedTemperature.Timestamp.Unix())
		}
	}

	for index, _ := range testSetsObjects {
		testSetsObjects[index].Delete()
	}
}

func TestAppend(t *testing.T) {
	var testObject SignedTemperatureSet = NewSignedTemperatureSet()
	var beginTime = time.Now()
	var prevTime = beginTime.Add(time.Duration(1) * time.Minute)

	testObject.Append(SignedTemperature{
		Temperature: Temperature.NewTemperature(-1),
		Timestamp:   beginTime,
	})
	testObject.Append(SignedTemperature{
		Temperature: Temperature.NewTemperature(-1),
		Timestamp:   prevTime,
	})

	if len(testObject) != 2 {
		t.Fatalf("testObject should have '2' elements, got '%d'", len(testObject))
	}
}
