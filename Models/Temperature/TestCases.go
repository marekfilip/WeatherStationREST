package Temperature

func getFloat64Slice() []float64 {
	return []float64{-2.123, -5.11, 2.009, 2.1, 5.11, 4.137, 10.0, 4.879, 5.300, 5.945, 6.512, -55.0, -54.999, 125, 124.75}
}

func getStringSlice() []string {
	return []string{"-2.123", "-5.11", "2.009", "2.1", "5.11", "4.137", "10.0", "4.879", "5.300", "5.945", "6.512", "-55.0", "-54.999", "125", "124.75"}
}

func GetUnparsableStringSlice() []string {
	return []string{"-1a", "a1.00", "as", ""}
}
