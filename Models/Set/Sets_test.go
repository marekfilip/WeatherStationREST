package Set

import (
	"testing"
	"time"
)

var testSetsObjects = GetSetsTestCases()
var testSets *Sets = nil
var startTime time.Time = time.Now()
var endTime time.Time = time.Now()

func TestNewSets(t *testing.T) {
	var zeroTime time.Time = time.Time{}

	if testSets != nil {
		t.Fatalf("Expected 'testSets' to be nil")
	}

	testSets = NewSets()

	if testSets == nil {
		t.Fatalf("Expected 'testSets' to not be")
	}

	if testSets.StartTime != zeroTime {
		t.Fatalf("Expected time to be '%v', got '%v'", zeroTime, testSets.StartTime)
	}

	if testSets.EndTime != zeroTime {
		t.Fatalf("Expected time to be '%v', got '%v'", zeroTime, testSets.EndTime)
	}

	if len(*testSets.List) != 0 {
		t.Fatalf("Expected length of sets list to be '0', got '%d'", len(*testSets.List))
	}
}

func TestAppend(t *testing.T) {
	if len(*testSets.List) != 0 {
		t.Fatalf("Expected length of sets list to be '0', got '%d'", len(*testSets.List))
	}

	for index, v := range testSetsObjects {
		endTime = startTime.Add(time.Duration(index) * time.Minute)
		v.SetCreated(endTime)

		testSets.Append(v)
	}

	if len(*testSets.List) != len(testSetsObjects) {
		t.Fatalf("Different count of objects in sets, expected '%d', got '%d'", len(testSetsObjects), len(*testSets.List))
	}

	if testSets.StartTime != startTime {
		t.Fatalf("Different startTime, expected '%d', got '%d'", startTime, testSets.StartTime)
	}

	if testSets.EndTime != endTime {
		t.Fatalf("Different endTime, expected '%d', got '%d'", endTime, testSets.EndTime)
	}
}

func TestLen(t *testing.T) {
	if len(*testSets.List) != testSets.Len() {
		t.Fatalf("Different count of objects in sets, expected '%d', got '%d'", len(*testSets.List), testSets.Len())
	}
}

func TestLess(t *testing.T) {
	if testSets.Less(0, 1) == false {
		t.Fatalf("Expected 0 to be less than 1")
	}

	if testSets.Less(1, 0) == true {
		t.Fatalf("Expected 0 to not be less than 1")
	}
}

func TestSwap(t *testing.T) {
	valueOf0 := (*testSets.List)[0].Brightness
	valueOf1 := (*testSets.List)[1].Brightness

	testSets.Swap(0, 1)

	if (*testSets.List)[0].Brightness != valueOf1 {
		t.Fatalf("After swap expected brightness 0 to be '%d', got '%d'", valueOf1, (*testSets.List)[0].Brightness)
	}

	if (*testSets.List)[1].Brightness != valueOf0 {
		t.Fatalf("After swap expected brightness 1 to be '%d', got '%d'", valueOf0, (*testSets.List)[1].Brightness)
	}
}

func TestFind(t *testing.T) {
	for index, _ := range testSetsObjects {
		testSetsObjects[index].Save()
	}

	sets, err := Find(startTime, endTime)
	if err != nil {
		t.Fatalf("Error occured:\n%s", err.Error())
	}

	if sets.StartTime.Unix() != testSetsObjects[0].Created.Unix() {
		t.Fatalf(
			"Different startTime, expected '%v', got '%v'",
			testSetsObjects[0].Created.Unix(),
			sets.StartTime.Unix(),
		)
	}

	if sets.EndTime.Unix() != testSetsObjects[len(testSetsObjects)-1].Created.Unix() {
		t.Fatalf("Different endTime, expected '%v', got '%v'", testSetsObjects[len(testSetsObjects)-1].Created.Unix(), sets.EndTime.Unix())
	}

	for i := 0; i < len(*sets.List)-1; i++ {
		if sets.Less(i, i+1) == false && (*sets.List)[i].Created != (*sets.List)[i+1].Created {
			t.Fatalf(`Seems that sets are not sorted '%d' is not less then '%d'
	%d. %+v
	%d. %+v\n`, i, i+1, i, (*sets.List)[i], i+1, (*sets.List)[i+1])
		}
	}

	for index, _ := range testSetsObjects {
		testSetsObjects[index].Delete()
	}
}
