package Temperature

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"
)

func Test(t *testing.T) {
	obj := Temperature{
		Timestamp: time.Unix(int64(1481720663), 0),
		Value:     big.NewFloat(13.123),
	}

	t.Logf("Before marshaling %+v\n", obj)
	a, err := json.Marshal(&obj)
	if err != nil {
		t.Errorf("Error when marshaling %s", err.Error())
		t.Fail()
	}
	t.Logf("Marshaled JSON %+v\n", string(a))

	var newObj Temperature
	err = json.Unmarshal(a, &newObj)
	if err != nil {
		t.Errorf("Error when unmarshaling %s", err.Error())
		t.Fail()
	}
	t.Logf("Unmarshaled object %+v\n", newObj)
}
