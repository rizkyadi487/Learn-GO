package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			helpMsg()
		} else {
			conversion(os.Args[1], os.Args[2], os.Args[3])
		}
	} else {
		errorMsg()
	}
}

func errorMsg() {
	println("Error: unrecognized or incomplete command line.")
	println("USAGE: go run main.go [value] [unit source] [unit conversion]")
	println("for more information go run main.go help")
}

func helpMsg() {
	println("list of command")
	println("cm	CentiMeter")
	println("m	Meter")
}

func conversion(v string, s string, c string) {
	switch s {
	case "cm":
		cmFormula(c, v)
	case "m":
		mFormula(c, v)
	}
}

func cmFormula(c string, v string) {
	i, _ := strconv.ParseFloat(v, 64)
	switch c {
	case "m":
		m := math.Pow10(2)
		fmt.Printf("%.0fm", m*i)
	default:
		println("mbalik gan")
	}
}

func mFormula(m string, v string) {
	i, _ := strconv.ParseFloat(v, 64)
	switch m {
	case "cm":
		cm := math.Pow10(2)
		fmt.Printf("%.0fcm", cm*i)
	default:
		println("mbalik gan")
	}
}
