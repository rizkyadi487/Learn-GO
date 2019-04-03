package main

import (
	"fmt"

	persegi "github.com/rizkyadi487/learngo/lib/persegi"
	"github.com/rizkyadi487/learngo/lib/segitiga"
)

func main() {

	var sisi = 9

	fmt.Printf("Luas Persegi adalah = %d\n", persegi.LuasPersegi(sisi))
	fmt.Printf("Keliling Persegi adalah = %d\n", persegi.KelilingPersegi(sisi))

	fmt.Printf("Luas Segitiga = %f\n", segitiga.LuasSegitiga(0.5, 9.4))

}
