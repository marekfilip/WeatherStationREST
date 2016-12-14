package Temperature

import (
	"encoding/json"

	util "filip/WeatherStationREST/Utilities"
)

type Temperatures map[int64]Temperature

func CreateFromEncryptedByteList(data []byte) (*Temperatures, error) {
	decrypted, err := util.Decrypt(data)
	if err != nil {
		return nil, err
	}

	var newObject Temperatures
	if err = json.Unmarshal(decrypted, &newObject); err != nil {
		return nil, err
	}

	return &newObject, nil
}

func (t *Temperatures) Encrypt() ([]byte, error) {
	return util.Encrypt(*t)
}

func (t Temperatures) GetAsByteList() ([]byte, error) {
	plainJsonByte, _ := json.Marshal(t)

	return plainJsonByte, nil
}

func (list *Temperatures) Append(objects map[int64]Temperature) {
	for timestamp, temp := range objects {
		(*list)[timestamp] = temp
	}
}

func (list *Temperatures) AppendNew(timestamp int64, celcius float32) {
	var temp Temperature = Temperature(celcius)

	var objects map[int64]Temperature = map[int64]Temperature{
		timestamp: temp,
	}

	list.Append(objects)
}
