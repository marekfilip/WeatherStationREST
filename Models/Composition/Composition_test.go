package Composition

import (
	"math/big"
	"testing"
	"time"

	temp "filip/WeatherStationREST/Models/Temperature"
)

func Test(t *testing.T) {
	var collectionObject Composition = Composition{
		Temperature: temp.Temperature{
			Timestamp: time.Unix(1481720663, 0),
			Value:     big.NewFloat(13.123),
		},
	}
	t.Logf("Before encryption: %+v\n", collectionObject)

	// Szyfrowanie
	encryptedCollectionObject, err := collectionObject.Encrypt()
	if err != nil {
		t.Errorf("Error when ecrypting %s", err.Error())
		t.Fail()
	}

	// Deszyfrowanie
	decryptedCollectionObject, err := CreateCompositionFromEncryptedByteList(encryptedCollectionObject)
	if err != nil {
		t.Errorf("Error when decrypting %s", err.Error())
		t.Fail()
	}
	t.Logf("After encryption: %+v\n", *decryptedCollectionObject)
}
