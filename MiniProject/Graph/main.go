package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 4,
		"b": 8,
		"c": 2,
		"d": 7,
	}

	a := len(m)
	fmt.Printf("Panjang Map : %d\n", a)
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
	graphHorizontal(m)
	graphVertical(m)
}

func graphHorizontal(m map[string]int) {
	for key, value := range m {
		fmt.Printf(key+" %d ", value)
		for i := 0; i < value; i++ {
			fmt.Print("-")
		}
		fmt.Println()
	}
}

//golang maps random order, Game is Hard
func graphVertical(m map[string]int) {
	maxValue := 0
	// for _, value := range m {
	// 	if maxValue < value {
	// 		maxValue = value
	// 	}
	// }

	keys, values := mapsToArray(m) //untuk menghidari radom order

	for i := 0; i < len(keys); i++ {
		if maxValue < values[i] {
			maxValue = values[i]
		}
	}

	fmt.Println("Max Value : ", maxValue)

	for i := maxValue; i > 0; i-- {
		fmt.Print(i)
		for j := 0; j < len(keys); j++ {
			if i <= values[j] {
				// fmt.Print(values[j])
				fmt.Print("| ")
			} else {
				// fmt.Print(values[j])
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	for i := 0; i < len(keys); i++ {
		fmt.Print(" ", values[i])
	}
}

func mapsToArray(m map[string]int) ([]string, []int) {

	var keys []string
	var values []int
	for key, value := range m {
		keys = append(keys, key)
		values = append(values, value)
	}

	// for index := 0; index < len(keys); index++ {
	// 	fmt.Print(keys[index])
	// 	fmt.Print(values[index])
	// }

	return keys, values
}
