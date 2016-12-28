package Composition

import (
	"encoding/json"
	"fmt"

	temp "filip/WeatherStationREST/Models/Temperature"
	util "filip/WeatherStationREST/Utilities"
)

type Composition struct {
	Temperature temp.Temperature `json:"temperature"`
}

func CreateCompositionFromEncryptedByteList(data []byte) (*Composition, error) {
	decrypted, err := util.Decrypt(data)
	if err != nil {
		return nil, err
	}

	var newObject Composition
	fmt.Printf("%+v\n", string(decrypted))
	if err = json.Unmarshal(decrypted, &newObject); err != nil {
		return nil, err
	}

	return &newObject, nil
}

func (c *Composition) Encrypt() ([]byte, error) {
	return util.Encrypt(*c)
}

func (c Composition) GetAsByteList() ([]byte, error) {
	plainJsonByte, _ := json.Marshal(&c)

	return plainJsonByte, nil
}
