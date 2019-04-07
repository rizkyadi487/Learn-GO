package main

import (
	"fmt"

	geometri "github.com/rizkyadi487/learngo/lib/geometri"
)

func main() {

	var a float64
	var b float64

	fmt.Printf("Masukkan Sisi Persegi : ")
	fmt.Scan(&a)

	luas, keliling := geometri.Persegi(a)

	fmt.Printf("Persegi dengan sisi %g mempunyai luas %g dan keliling %g\n", a, luas, keliling)
	blank()
	fmt.Printf("Masukkan panjang dari persegi panjang :")
	fmt.Scan(&a)
	fmt.Printf("Masukkan lebar dari persegi panjang :")
	fmt.Scan(&b)
	fmt.Printf("Persegi panjang dengan panjang %g dan lebar %g\n", a, b)

	luas, keliling = geometri.PersegiPanjang(a, b)

	fmt.Printf("Luas Persegi Panjang : %g\n", luas)
	fmt.Printf("Keliling Persegi Panjang : %g\n", keliling)
	blank()
	fmt.Printf("Masukkan jari jari lingkaran : ")
	fmt.Scan(&a)
	luas, keliling = geometri.Lingkaran(a)
	fmt.Printf("Keliling lingkaran = %.3f\n", keliling)
	fmt.Printf("Luas lingkaran = %.3f\n", luas)
}
func blank() {
	fmt.Println("\n======================================")
}
