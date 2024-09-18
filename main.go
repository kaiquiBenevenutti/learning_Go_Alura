package main

import "fmt"

func main() {
	nome := "Kaiqui"
	versao := 1.0
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão", versao)

	fmt.Println("")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando) // o & e para mostrar o caminho para onde a resposta deve ser salva, nesse caso o comando

	fmt.Println("O comando escolhido foi", comando)
}
