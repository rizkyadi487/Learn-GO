package main

import "runtime"

func main() {
	jumlahCPU := runtime.NumCPU()
	println("Jumlah CPU = ", jumlahCPU)
}
