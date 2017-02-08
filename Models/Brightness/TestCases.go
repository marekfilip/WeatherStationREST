package Brightness

func getTestCases() []struct {
	temperature float64
	timestamp   int64
	object      *Brightness
} {
	return []struct {
		temperature float64
		timestamp   int64
		object      *Brightness
	}{
		{temperature: 970, timestamp: 1482192000},  // Wed, 20 Dec 2016 00:00:00 GMT
		{temperature: 953, timestamp: 1482278400},  // Wed, 21 Dec 2016 00:00:00 GMT
		{temperature: 900, timestamp: 1482368400},  // Thu, 22 Dec 2016 01:00:00 GMT
		{temperature: 568, timestamp: 1482454980},  // Fri, 23 Dec 2016 01:03:00 GMT
		{temperature: 498, timestamp: 1482587711},  // Sat, 24 Dec 2016 13:55:11 GMT
		{temperature: 123, timestamp: 1482670271},  // Sun, 25 Dec 2016 12:51:11 GMT
		{temperature: 357, timestamp: 1482772509},  // Mon, 26 Dec 2016 17:15:09 GMT
		{temperature: 596, timestamp: 1482829200},  // Tue, 27 Dec 2016 09:00:00 GMT
		{temperature: 875, timestamp: 1482832800},  // Tue, 27 Dec 2016 10:00:00 GMT
		{temperature: 598, timestamp: 1482836400},  // Tue, 27 Dec 2016 11:00:00 GMT
		{temperature: 1000, timestamp: 1482840000}, // Tue, 27 Dec 2016 12:00:00 GMT
	}
}

func getMinTimestamp() int64 {
	var min int64 = 0

	for _, c := range getTestCases() {
		if min == 0 || c.timestamp < min {
			min = c.timestamp
		}
	}

	return min
}

func getMaxTimestamp() int64 {
	var max int64 = 0

	for _, c := range getTestCases() {
		if max == 0 || c.timestamp > max {
			max = c.timestamp
		}
	}

	return max
}
