package bangunruang

func Volume_Balok(p, l, t int) int {
	return p * l * t
}
func Luas_Permukaan_Balok(p, l, t int) int {
	return 2 * (p*l + p*t + l*t)
}
func Panjang_Balok(p, l, t int) int {
	return (p * l * t) / l / t
}
