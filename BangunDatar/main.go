package main

import (
	"fmt"

	persegi "github.com/rizkyadi487/learngo/lib/persegi"
	segitiga "github.com/rizkyadi487/learngo/lib/segitiga"
)

func main() {

	var sisi int

	fmt.Printf("Masukkan Sisi Persegi : ")
	fmt.Scan(&sisi)
	fmt.Printf("Luas Persegi adalah = %d\n", persegi.Luas(sisi))
	fmt.Printf("Keliling Persegi adalah = %d\n", persegi.Keliling(sisi))

	fmt.Printf("Luas Segitiga = %f\n", segitiga.Luas(0.5, 9.4))
	fmt.Printf("Keliling Segitiga = %f\n", segitiga.Keliling(3, 2, 1))
}
