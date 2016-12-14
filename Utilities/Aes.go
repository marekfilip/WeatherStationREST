package Utilities

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"filip/WeatherStationREST/Config"
	"filip/WeatherStationREST/Models"
)

func Encrypt(object Entities.Byteable) ([]byte, error) {
	data, err := object.GetAsByteList()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(Config.GetSecret())
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

func Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(Config.GetSecret())
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("Ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	decodedString, err := base64.StdEncoding.DecodeString(string(data))

	if err != nil {
		return nil, err
	}

	return decodedString, nil
}
