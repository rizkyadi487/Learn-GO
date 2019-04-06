package geometri

//Persegi Mempunyai masukkan sisi(float64)
//Kembalian (luas float64, keliling float64)
func Persegi(sisi float64) (luas float64, keliling float64) {
	luas = sisi * sisi
	keliling = sisi * 4
	return luas, keliling
}

//PersegiPanjang Mempunyai masukkan panjang(float64) dan lebar(float64)
//Kembalian (luas float64, keliling float64)
func PersegiPanjang(panjang float64, lebar float64) (luas float64, keliling float64) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return luas, keliling
}
