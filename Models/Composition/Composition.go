package Composition

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	conf "filip/WeatherStationREST/Config"
	temp "filip/WeatherStationREST/Models/Temperature"
)

type Composition struct {
	Temperature temp.Temperature `json:"temperature"`
}

func CreateFromEncryptedBytes(data []byte) (*Composition, error) {
	var newObject Composition

	err := newObject.decryptTo(data)
	if err != nil {
		return nil, err
	}

	return &newObject, nil
}

func (c *Composition) Encrypt() ([]byte, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(conf.GetSecret())
	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(data)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return ciphertext, nil
}

func (c *Composition) decryptTo(data []byte) error {
	block, err := aes.NewCipher(conf.GetSecret())
	if err != nil {
		return err
	}

	if len(data) < aes.BlockSize {
		return errors.New("Ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	decodedString, err := base64.StdEncoding.DecodeString(string(data))

	if err != nil {
		return err
	}

	if err = json.Unmarshal(decodedString, c); err != nil {
		return err
	}

	return nil
}
