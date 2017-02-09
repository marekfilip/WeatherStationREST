package Brightness

func getUint16Slice() []uint16 {
	return []uint16{970, 953, 900, 568, 498, 123, 357, 596, 875, 598, 1000}
}

func getStringSlice() []string {
	return []string{"970", "953", "900", "568", "498", "123", "357", "596", "875", "598", "1000"}
}

func GetUnparsableStringSlice() []string {
	return []string{"-1", "1.00", "70000", ""}
}
