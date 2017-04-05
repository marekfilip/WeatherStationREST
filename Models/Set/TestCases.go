package Set

import (
	"filip/WeatherStationREST/Models/Brightness"
	"filip/WeatherStationREST/Models/Temperature"
)

func getMapSetsTestCases() []map[string]string {
	return []map[string]string{
		map[string]string{
			"B": "970",
			"T": "-2.123",
		},
		map[string]string{
			"B": "953",
			"T": "-5.11",
		},
		map[string]string{
			"B": "900",
			"T": "2.009",
		},
		map[string]string{
			"B": "568",
			"T": "2.1",
		},
		map[string]string{
			"B": "498",
			"T": "5.11",
		},
		map[string]string{
			"B": "123",
			"T": "4.137",
		},
		map[string]string{
			"B": "357",
			"T": "10.0",
		},
		map[string]string{
			"B": "596",
			"T": "4.879",
		},
		map[string]string{
			"B": "875",
			"T": "5.300",
		},
		map[string]string{
			"B": "598",
			"T": "5.945",
		},
		map[string]string{
			"B": "1000",
			"T": "6.512",
		},
	}
}

func GetSetsTestCases() []Set {
	return []Set{
		Set{
			Brightness:  Brightness.NewBrightness(970),
			Temperature: Temperature.NewTemperature(-2.123),
		},
		Set{
			Brightness:  Brightness.NewBrightness(953),
			Temperature: Temperature.NewTemperature(-5.11),
		},
		Set{
			Brightness:  Brightness.NewBrightness(900),
			Temperature: Temperature.NewTemperature(2.009),
		},
		Set{
			Brightness:  Brightness.NewBrightness(568),
			Temperature: Temperature.NewTemperature(2.1),
		},
		Set{
			Brightness:  Brightness.NewBrightness(498),
			Temperature: Temperature.NewTemperature(5.11),
		},
		Set{
			Brightness:  Brightness.NewBrightness(123),
			Temperature: Temperature.NewTemperature(4.137),
		},
		Set{
			Brightness:  Brightness.NewBrightness(357),
			Temperature: Temperature.NewTemperature(10.0),
		},
		Set{
			Brightness:  Brightness.NewBrightness(596),
			Temperature: Temperature.NewTemperature(4.879),
		},
		Set{
			Brightness:  Brightness.NewBrightness(875),
			Temperature: Temperature.NewTemperature(5.300),
		},
		Set{
			Brightness:  Brightness.NewBrightness(598),
			Temperature: Temperature.NewTemperature(5.945),
		},
		Set{
			Brightness:  Brightness.NewBrightness(1000),
			Temperature: Temperature.NewTemperature(6.512),
		},
	}
}
