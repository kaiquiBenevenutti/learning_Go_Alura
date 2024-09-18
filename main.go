package main

import (
	"fmt"
	"os"
)

func main() {
	bemVindo()
	fmt.Println("")

	exibirMenu()
	comando := lerComando()
	fmt.Println("")

	switch comando {
	case 1:
		fmt.Println("Monitoramento Iniciado.")
	case 2:
		fmt.Println("Exibindo logs.")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0) //para informar ao computador que saiu com sucesso, sem erros
	default:
		fmt.Println("Tente novamente.")
		os.Exit(-1) //para informar ao computador que algo deu errado
	}
}

func bemVindo() {
	nome := "Kaiqui"
	versao := 1.2
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão", versao)
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // o & e para mostrar o caminho para onde a resposta deve ser salva, nesse caso o comando
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
