package Brightness

import (
	"testing"
)

var brightnessTestObject BrightnessSet

func TestAppendBrightness(t *testing.T) {
	if len(brightnessTestObject) != 0 {
		t.Fatalf("Expected to have 0 elements, got %d", len(brightnessTestObject))
	}

	for index, value := range getUint16Slice() {
		t.Logf("Appending object '%d'", index)
		brightnessTestObject.Append(Brightness(value))
	}

	if len(brightnessTestObject) != len(getUint16Slice()) {
		t.Fatalf("Expected to have '%d' elements, got %d", len(getUint16Slice()), len(brightnessTestObject))
	}
}
