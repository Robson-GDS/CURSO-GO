package main

import "fmt"

type ID int // criando type

var usandoTypeId ID = 1 // add type em uma variavel

// usanto printf é possivel usar %T que apresenta o tipo da variavel
// %v apresenta o valor

func main() {
	fmt.Printf("%T", usandoTypeId)
}
