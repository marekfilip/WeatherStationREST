package main

import (
	"log"

	conf "filip/WeatherStationREST/Config"

	"github.com/maxwellhealth/bongo"
	//coll "filip/WeatherStationREST/Models/Collection"
	//temp "filip/WeatherStationREST/Models/Temperature"
)

func main() {
	/*var collectionObject coll.Collection = coll.Collection{
		Temperatures: temp.Temperatures{
			1481720663: 13,
			1481720664: 13.123,
		},
	}

	// Szyfrowanie
	log.Printf("Przed szyfrowaniem %T %+v\n\n", collectionObject, collectionObject)
	encryptedCollectionObject, err := collectionObject.Encrypt()
	if err != nil {
		panic(err)
	}

	// Deszyfrowanie
	newObj, err := coll.CreateFromEncryptedByteList(encryptedCollectionObject)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Po deszyfrowaniu %T %+v\n\n", *newObj, *newObj)*/
}
