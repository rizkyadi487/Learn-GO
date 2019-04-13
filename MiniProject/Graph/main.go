package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 3711,
		"b": 2138,
		"c": 1908,
		"d": 912,
	}

	a := len(m)
	fmt.Printf("Panjang Map : %d\n", a)
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
