package Config

import (
	"os"

	"github.com/maxwellhealth/bongo"
)

var environment = os.Getenv("WEATHER_STATION_ENV")

const (
	productionSetCollectionName = "set"
	testingSetCollectionName    = "setTest"
)

var config = &bongo.Config{
	ConnectionString: "mongodb://USER:PASSWORD@HOST:PORT/DB_NAME",
	Database:         "DB_NAME",
}

// GetBongoConfig gets the configuration for Bongo instance
func GetBongoConfig() *bongo.Config {
	return config
}

// GetSetCollectionName gets the collection name based on environment
func GetSetCollectionName() string {
	switch environment {
	case "prod":
		return productionSetCollectionName
	}

	return testingSetCollectionName
}
