package Set

func getTestCases() []struct {
	temperature float64
	timestamp   int64
	object      *Set
	encrypted   []byte
} {
	return []struct {
		temperature float64
		timestamp   int64
		object      *Set
		encrypted   []byte
	}{
		{temperature: -2.123, timestamp: 1482192000}, // Wed, 20 Dec 2016 00:00:00 GMT
		{temperature: -5.11, timestamp: 1482278400},  // Wed, 21 Dec 2016 00:00:00 GMT
		{temperature: 2.009, timestamp: 1482368400},  // Thu, 22 Dec 2016 01:00:00 GMT
		{temperature: 2.1, timestamp: 1482454980},    // Fri, 23 Dec 2016 01:03:00 GMT
		{temperature: 5.11, timestamp: 1482587711},   // Sat, 24 Dec 2016 13:55:11 GMT
		{temperature: 4.137, timestamp: 1482670271},  // Sun, 25 Dec 2016 12:51:11 GMT
		{temperature: 10.0, timestamp: 1482772509},   // Mon, 26 Dec 2016 17:15:09 GMT
		{temperature: 4.879, timestamp: 1482829200},  // Tue, 27 Dec 2016 09:00:00 GMT
		{temperature: 5.300, timestamp: 1482832800},  // Tue, 27 Dec 2016 10:00:00 GMT
		{temperature: 5.945, timestamp: 1482836400},  // Tue, 27 Dec 2016 11:00:00 GMT
		{temperature: 6.512, timestamp: 1482840000},  // Tue, 27 Dec 2016 12:00:00 GMT
	}
}
