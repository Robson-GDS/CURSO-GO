package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Name string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Name  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Name)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	robson := Cliente{
		Name:  "Robson",
		Idade: 32,
		Ativo: true,
	}

	Desativacao(robson)
}
