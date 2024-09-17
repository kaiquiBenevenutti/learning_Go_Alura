package main

import (
	"fmt"
	"strings"
)

func main() {

	var resposta string

	fmt.Print("Digite o seu nome: ")
	_, err := fmt.Scanln(&resposta)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	fmt.Println("Ola", resposta, "como voce esta?")

	var resposta2 string
	resposta2 = strings.ToLower(resposta2)

	for {
		_, err2 := fmt.Scanln(&resposta2)

		if err2 != nil {
			fmt.Println("Erro ao ler a entrada:", err2)
			return
		}

		if resposta2 == "bem" {
			fmt.Println("Que otimo!")
			break
		} else if resposta2 == "mal" {
			fmt.Println("Ah, que pena.")
			break
		} else {
			fmt.Println("Não entendi sua resposta, responda com 'bem' ou 'mal'.")
			fmt.Println(resposta, "como voce esta?")
		}
	}
}
