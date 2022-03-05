package bangunruang

import "math"

func Volume_Kubus(s int) int {
	return int(math.Pow(float64(s), 3))
}
func Luas_Permukaan(s int) int {
	return 6 * (int(math.Pow(float64(s), 2)))
}
func Keliling_Kubus(s int) int {
	return 12 * (int(math.Pow(float64(s), 2)))
}
func Luas_Sisi(s int) int {
	return int(math.Pow(float64(s), 2))
}
