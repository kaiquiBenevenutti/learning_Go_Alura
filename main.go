package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	bemVindo()
	fmt.Println("")

	for {
		exibirMenu()
		comando := lerComando()
		fmt.Println("")

		switch comando {
		case 1:
			iniciarMonitoramento()
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
}

func bemVindo() {
	nome := "Kaiqui"
	versao := 1.9
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			resp, err := http.Get(site)
			if err != nil {
				fmt.Println("Ocorreu um erro:", err)
				return
			}
			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "foi carregado com sucesso!")
			} else {
				fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
			}
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func lerSitesDoArquivo() []string {
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	linha, err := leitor.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(linha)
	return []string{"oi"}
}
