package main

import "fmt"

func main() {
	bemVindo()
	fmt.Println("")

	exibirMenu()
	comando := lerComando()
	fmt.Println("")
	
	definirEscolha(comando)
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

func lerComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido) // o & e para mostrar o caminho para onde a resposta deve ser salva, nesse caso o comando
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func definirEscolha(comando int) {
	switch comando {
	case 1:
		fmt.Println("Monitoramento Iniciado.")
	case 2:
		fmt.Println("Exibindo logs.")
	case 0:
		fmt.Println("Obrigado, até a proxima!")
	default:
		fmt.Println("Tente novamente.")
	}
}
