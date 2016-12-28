package Temperature

import (
	"math/big"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var temperaturesTestObject Temperatures

func TestAppendTemperatures(t *testing.T) {
	if len(temperaturesTestObject) != 0 {
		t.Fatalf("Expected to have 0 elements, got %d", len(temperaturesTestObject))
	}

	var temp Temperature
	for index, c := range getTestCases() {
		temp = Temperature{
			Timestamp: time.Unix(c.timestamp, 0),
			Value:     big.NewFloat(c.temperature),
		}

		t.Logf("Appending object '%d'", index)
		temperaturesTestObject.Append(temp)
	}

	if len(temperaturesTestObject) != len(getTestCases()) {
		t.Fatalf("Expected to have '%d' elements, got %d", len(getTestCases()), len(temperaturesTestObject))
	}
}

func TestLen(t *testing.T) {
	if len(temperaturesTestObject) != temperaturesTestObject.Len() {
		t.Fatalf("Expected to have '%d' elements, got %d", len(temperaturesTestObject), temperaturesTestObject.Len())
	}
}

func TestSwap(t *testing.T) {
	if temperaturesTestObject.Len() <= 0 {
		t.Fatalf("Can not test swap if there is no elements")
	}

	first := temperaturesTestObject[0]
	second := temperaturesTestObject[1]

	temperaturesTestObject.Swap(0, 1)

	if temperaturesTestObject[0].Timestamp == first.Timestamp && temperaturesTestObject[0].Value == first.Value {
		t.Errorf("It seems that objects did not swap, expected '%v', got '%v' and expected '%v', got '%v'", second.Timestamp, first.Timestamp, second.Value, first.Value)
	}
}

func TestSort(t *testing.T) {
	if temperaturesTestObject.Len() <= 0 {
		t.Fatalf("Can not test sort if there is no elements")
	}

Loop:
	for {
		t.Logf("Checking that objects are shuffled")
		for i := 0; i < temperaturesTestObject.Len()-2; i++ {
			if temperaturesTestObject[i].Timestamp.After(temperaturesTestObject[i+1].Timestamp) {
				break Loop
			}
		}

		// Shuffle three times
		for i := 0; i < 3; i++ {
			t.Logf("Shuffling '%d'", i)
			temperaturesTestObject.Swap(rand.Intn(temperaturesTestObject.Len()), rand.Intn(temperaturesTestObject.Len()))
		}
	}

	t.Logf("Sorting and checking")
	sort.Sort(temperaturesTestObject)
	for i := 1; i < temperaturesTestObject.Len(); i++ {
		if temperaturesTestObject.Less(i-1, i) == false {
			t.Fatalf("Expected that '%v' is lower then '%v'", temperaturesTestObject[i-1].Timestamp, temperaturesTestObject[i].Timestamp)
		}
	}
}

func TestLess(t *testing.T) {
	for i := 1; i < temperaturesTestObject.Len(); i++ {
		if temperaturesTestObject.Less(i-1, i) == false {
			t.Fatalf("Expected that '%v' is lower then '%v'", temperaturesTestObject[i-1].Timestamp, temperaturesTestObject[i].Timestamp)
		}
	}
}

func TestFindTemperatures(t *testing.T) {
	var err error
	var newTemperatures Temperatures

	t.Logf("Saving objects")
	for i, _ := range temperaturesTestObject {
		err = temperaturesTestObject[i].Save()
		if err != nil {
			t.Errorf("Error when saving object: %s", err.Error())
		}
	}

	t.Logf("Finding objects")
	err = newTemperatures.Find(time.Unix(getMinTimestamp(), 0), time.Unix(getMaxTimestamp(), 0))
	if err != nil {
		t.Fatalf("Error when finding objects: %s", err.Error())
	}

	t.Logf("Sorting both structures")
	sort.Sort(temperaturesTestObject)
	sort.Sort(newTemperatures)

	if temperaturesTestObject.Len() != newTemperatures.Len() {
		t.Fatalf("Found objects length different then saved, expected '%d', got '%d'", temperaturesTestObject.Len(), newTemperatures.Len())
	}

	t.Logf("Comparing them")
	for index, c := range temperaturesTestObject {
		if c.Timestamp != newTemperatures[index].Timestamp || c.Value.Cmp(newTemperatures[index].Value) != 0 {
			t.Fatalf("Found different objects '%+v' and '%+v'", c, newTemperatures[index])
		}
	}

	t.Logf("Removing objects")
	for i, _ := range temperaturesTestObject {
		err = temperaturesTestObject[i].Delete()
		if err != nil {
			t.Errorf("Error when deleting object: %s", err.Error())
		}
	}
}
