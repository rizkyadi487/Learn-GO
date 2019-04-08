package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		cm, _ := strconv.ParseInt(os.Args[1], 10, 64)
		println("Meter = ", os.Args[1])
		println("Centi Meter = ", cm*100)
	} else {
		println("wrong format")
		println("example : go run main.go 90")
	}
}
