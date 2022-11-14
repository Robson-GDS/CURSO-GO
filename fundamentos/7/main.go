package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{"Wesley": 1000, "Joao": 2000, "Maria": 3000}
	fmt.Println(salarios["Joao"])

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}
}
