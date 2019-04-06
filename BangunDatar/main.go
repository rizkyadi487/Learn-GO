package main

import (
	"fmt"

	geometri "github.com/rizkyadi487/learngo/lib/geometri"
)

func main() {

	var sisi float64
	var panjang float64
	var lebar float64

	fmt.Printf("Masukkan Sisi Persegi : ")
	fmt.Scan(&sisi)
	luas, keliling := geometri.Persegi(sisi)
	fmt.Printf("Persegi dengan sisi %g mempunyai luas %g dan keliling %g\n", sisi, luas, keliling)

	fmt.Printf("Masukkan panjang dari persegi panjang :")
	fmt.Scan(&panjang)
	fmt.Printf("Masukkan lebar dari persegi panjang :")
	fmt.Scan(&lebar)
	fmt.Printf("Persegi panjang dengan panjang %g dan lebar %g\n", panjang, lebar)
	luas, keliling = geometri.PersegiPanjang(panjang, lebar)
	fmt.Printf("Luas Persegi Panjang : %g\n", luas)
	fmt.Printf("Keliling Persegi Panjang : %g\n", keliling)

}
