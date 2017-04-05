package Signed

import (
	"testing"
	"time"

	bri "filip/WeatherStationREST/Models/Brightness"
	"filip/WeatherStationREST/Models/Set"
)

func TestFind(t *testing.T) {
	var testSetsObjects = Set.GetSetsTestCases()

	for index := range testSetsObjects {
		testSetsObjects[index].Save()
	}

	signedBrightnesss, err := Find(
		testSetsObjects[0].Created,
		testSetsObjects[len(testSetsObjects)-1].Created,
	)
	if err != nil {
		t.Fatalf("Error occured:\n%s", err.Error())
	}

outerLoop:
	for _, testSetObject := range testSetsObjects {
		for index, signedBrightness := range signedBrightnesss {
			if signedBrightness.Timestamp.Unix() == testSetObject.Created.Unix() {
				signedBrightnesss.Delete(uint(index))
				continue outerLoop
			}

			t.Fatalf("Could not find object with timestamp '%v'", signedBrightness.Timestamp)
		}
	}

	for index := range testSetsObjects {
		testSetsObjects[index].Delete()
	}
}

func TestAppend(t *testing.T) {
	var testObject = NewBrightnessSet()
	var beginTime = time.Now()
	var prevTime = beginTime.Add(time.Duration(1) * time.Minute)

	testObject.Append(Brightness{
		Brightness: bri.NewBrightness(100),
		Timestamp:  beginTime,
	})
	testObject.Append(Brightness{
		Brightness: bri.NewBrightness(100),
		Timestamp:  prevTime,
	})

	if len(testObject) != 2 {
		t.Fatalf("testObject should have '2' elements, got '%d'", len(testObject))
	}
}
