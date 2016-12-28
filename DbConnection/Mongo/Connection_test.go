package Mongo

import (
	"testing"

	"github.com/maxwellhealth/bongo"
)

func TestGetConnection(t *testing.T) {
	var firstConnection, nextConnection *bongo.Connection
	var err error

	firstConnection, err = GetConnection()
	if err != nil {
		t.Fatalf("Got error on connection: '%s'", err.Error())
	}

	for i := 0; i < 10; i++ {
		nextConnection, err = GetConnection()
		if err != nil {
			t.Fatalf("Got error on connection: '%s'", err.Error())
		}

		if firstConnection != nextConnection {
			t.Errorf("Pointer to connection changed on '%d' try, expected '%s', now got '%s'", i, firstConnection, nextConnection)
		}
	}
}
