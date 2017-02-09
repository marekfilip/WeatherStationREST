package Mongo

import (
	"testing"

	"github.com/maxwellhealth/bongo"
)

func TestGetConnection(t *testing.T) {
	var firstConnection, nextConnection *bongo.Connection

	firstConnection = GetConnection()
	for i := 0; i < 10; i++ {
		nextConnection = GetConnection()

		if firstConnection != nextConnection {
			t.Errorf("Pointer to connection changed on '%d' try, expected '%s', now got '%s'", i, firstConnection, nextConnection)
		}
	}
}
