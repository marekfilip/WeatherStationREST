package Composition

import (
	"math/big"
	"testing"
	"time"

	temp "filip/WeatherStationREST/Models/Temperature"
)

var compositionTestCases = getTestCases()

func TestCreateComposition(t *testing.T) {
	var value float64

	for index, c := range compositionTestCases {
		compositionTestCases[index].object = &Composition{
			Temperature: temp.Temperature{
				Timestamp: time.Unix(c.timestamp, 0),
				Value:     big.NewFloat(c.temperature),
			},
		}
	}

	for index, c := range compositionTestCases {
		t.Logf("Creating object '%d'", index)

		if c.object.Temperature.Timestamp.Unix() != c.timestamp {
			t.Fatalf("Expected timestamp is '%d' got '%d'", index, c.timestamp, c.object.Temperature.Timestamp.Unix())
		}

		value, _ = c.object.Temperature.Value.Float64()
		if value != c.temperature {
			t.Fatalf("Expected value is '%f' got '%f'", index, c.temperature, value)
		}
	}
}

func TestEncrypt(t *testing.T) {
	var err error

	for index, c := range compositionTestCases {
		t.Logf("Encrypting object '%d'", index)

		compositionTestCases[index].encrypted, err = c.object.Encrypt()
		if err != nil {
			t.Fatalf("Error when ecrypting: %s", err.Error())
		}
	}
}

func TestDecrypt(t *testing.T) {
	var tempComposition *Composition
	var err error

	for index, c := range compositionTestCases {
		t.Logf("Decrypting object '%d'", index)

		tempComposition, err = CreateFromEncryptedBytes(c.encrypted)
		if err != nil {
			t.Fatalf("Error when decrypting: %s", err.Error())
		}

		if tempComposition.Temperature.Timestamp != c.object.Temperature.Timestamp || tempComposition.Temperature.Value.Cmp(c.object.Temperature.Value) != 0 {
			t.Fatalf("Decrypt failed, expected '%v', got '%v'", c.object, tempComposition)
		}
	}
}
