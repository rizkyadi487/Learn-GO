package geometri

//Persegi Mempunyai masukkan sisi(float64)
//Kembalian (luas float64, keliling float64)

const phi float64 = 3.14159265358979323846

func Persegi(s float64) (luas float64, keliling float64) {
	luas = s * s
	keliling = s * 4
	return luas, keliling
}

//PersegiPanjang Mempunyai masukkan panjang(float64) dan lebar(float64)
//Kembalian (luas float64, keliling float64)
func PersegiPanjang(p float64, l float64) (luas float64, keliling float64) {
	luas = p * l
	keliling = 2 * (p + l)
	return luas, keliling
}

func Lingkaran(r float64) (luas float64, keliling float64) {
	luas = phi * r * r
	keliling = 2 * phi * r
	return luas, keliling
}
