package Brightness

import "testing"

func TestParseString(t *testing.T) {
	var b *Brightness
	var err error
	var expected []uint16 = getUint16Slice()

	for index, value := range getStringSlice() {
		t.Logf("Parsing '%d' string", index)

		b, err = ParseString(value)
		if err != nil {
			t.Fatalf("Error occured\n%s", err.Error())
		}

		if uint16(*b) != expected[index] {
			t.Fatalf("Error when parsing, expected '%d' got '%d'", expected[index], *b)
		}
	}
}

func TestUnparsableParseString(t *testing.T) {
	var b *Brightness
	var err error

	for index, value := range GetUnparsableStringSlice() {
		t.Logf("Parsing '%d' string", index)

		b, err = ParseString(value)
		if err == nil {
			t.Fatalf("Expected error when parsing '%s', got '%d'", value, *b)
		}
	}
}
