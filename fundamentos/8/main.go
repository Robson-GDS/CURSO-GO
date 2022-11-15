package main

import "fmt"

func main() {
	fmt.Println(sum(70, 2))
}

func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}
