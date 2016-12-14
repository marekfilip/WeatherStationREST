package Entities

type Encryptable interface {
	Encrypt() ([]byte, error)
}

type Byteable interface {
	GetAsByteList() ([]byte, error)
}
