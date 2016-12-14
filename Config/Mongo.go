package Config

import "github.com/maxwellhealth/bongo"

var config *bongo.Config = &bongo.Config{
	ConnectionString: "mongodb://filip:magmagmag90@ds139187.mlab.com:39187",
	Database:         "weather-station",
}

func GetMongoConfig() *bongo.Config {
	return config
}
