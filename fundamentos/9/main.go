package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3, 4, 5, 102, 180404, 700))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
