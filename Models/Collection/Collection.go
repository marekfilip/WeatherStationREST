package Collection

import (
	"encoding/json"

	temp "filip/WeatherStationREST/Models/Temperature"
	util "filip/WeatherStationREST/Utilities"
)

type Collection struct {
	Temperatures temp.Temperatures
}

func CreateFromEncryptedByteList(data []byte) (*Collection, error) {
	decrypted, err := util.Decrypt(data)
	if err != nil {
		return nil, err
	}

	var newObject Collection
	if err = json.Unmarshal(decrypted, &newObject); err != nil {
		return nil, err
	}

	return &newObject, nil
}

func (c *Collection) Encrypt() ([]byte, error) {
	return util.Encrypt(*c)
}

func (c Collection) GetAsByteList() ([]byte, error) {
	plainJsonByte, _ := json.Marshal(c)

	return plainJsonByte, nil
}
