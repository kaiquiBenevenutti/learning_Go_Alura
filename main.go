package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Kaiqui"
	nome2 := "Ju"
	var idadeKaiqui int = 18
	idadeJu := 78
	var numero float32 = 1.5
	numero2 := 2.55

	fmt.Println("Nome:", nome)
	fmt.Println("Nome 2:", nome2)
	fmt.Println("Idade Kaiqui:", idadeKaiqui)
	fmt.Println("Idade Ju:", idadeJu)
	fmt.Println("Número 1:", numero)
	fmt.Println("Número 2:", numero2)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}
