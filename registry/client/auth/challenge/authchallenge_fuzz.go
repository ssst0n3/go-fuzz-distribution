package challenge

func Fuzz_parseValueAndParams(data []byte) int {
	parseValueAndParams(string(data))
	return 0
}
