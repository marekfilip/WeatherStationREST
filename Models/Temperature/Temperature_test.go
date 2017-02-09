package Temperature

import (
	"math/big"
	"testing"
)

func TestNewTemperature(t *testing.T) {
	var b *Temperature
	var expected []float64 = getFloat64Slice()

	for index, value := range getFloat64Slice() {
		t.Logf("Creating new temperature '%d'", index)

		b = NewTemperature(value)
		if v, _ := b.Float64(); v != expected[index] {
			t.Fatalf("Error when parsing, expected '%f' got '%f'", expected[index], v)
		} else {
			t.Logf("Parsed, expected '%f' got '%f'", expected[index], v)
		}
	}
}

func TestFloat64(t *testing.T) {
	var b *Temperature
	var floatValue float64
	var accuracy big.Accuracy
	var expected []float64 = getFloat64Slice()

	for index, value := range getFloat64Slice() {
		t.Logf("Creating new temperature '%d'", index)

		b = NewTemperature(value)
		floatValue, accuracy = b.Float64()
		if floatValue != expected[index] {
			t.Fatalf("Error when parsing, expected '%f' got '%f'", expected[index], floatValue)
		} else if accuracy != big.Exact {
			t.Fatalf("Error when parsing, expected '%v' got '%v'", accuracy, big.Exact)
		} else {
			t.Logf("Parsed, expected '%f' got '%f'", expected[index], floatValue)
		}
	}
}

func TestParseString(t *testing.T) {
	var b *Temperature
	var err error
	var expected []float64 = getFloat64Slice()

	for index, value := range getStringSlice() {
		t.Logf("Parsing '%d' string", index)

		b, err = ParseString(value)
		if err != nil {
			t.Fatalf("Error occured\n%s", err.Error())
		}

		if v, _ := b.Float64(); v != expected[index] {
			t.Fatalf("Error when parsing, expected '%f' got '%f'", expected[index], v)
		} else {
			t.Logf("Parsed, expected '%f' got '%f'", expected[index], v)
		}
	}
}

func TestUnparsableParseString(t *testing.T) {
	var b *Temperature
	var err error

	for index, value := range GetUnparsableStringSlice() {
		t.Logf("Parsing '%d' string", index)

		b, err = ParseString(value)
		if err == nil {
			t.Fatalf("Expected error when parsing '%s', got '%d'", value, *b)
		}
	}
}
