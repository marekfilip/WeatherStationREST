package Mongo

func GetConnection() {
	connection, err := bongo.Connect(conf.GetMongoConfig())

	if err != nil {
		log.Fatalln(err.Error())
	}
}
