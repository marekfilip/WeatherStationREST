package Temperature

import (
	"testing"
)

var temperaturesTestObject TemperatureSet

func TestAppendTemperatures(t *testing.T) {
	var temp *Temperature

	if len(temperaturesTestObject) != 0 {
		t.Fatalf("Expected to have 0 elements, got %d", len(temperaturesTestObject))
	}

	for index, value := range getFloat64Slice() {
		t.Logf("Appending object '%d'", index)
		temp = NewTemperature(value)
		temperaturesTestObject.Append(*temp)
	}

	if len(temperaturesTestObject) != len(getFloat64Slice()) {
		t.Fatalf("Expected to have '%d' elements, got %d", len(getFloat64Slice()), len(temperaturesTestObject))
	}
}
