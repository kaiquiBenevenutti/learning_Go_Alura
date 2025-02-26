package main

import (
	"Primeiro_Projeto/clientes"
	"Primeiro_Projeto/contas"
	"fmt"
)

func main() {
	contaNova := contas.ContaCorrente{}
	contaNova.Titular = clientes.Titular{"Nova conta", "767857858787587", "Teste"}
	contaNova.NumeroConta = 232443
	contaNova.NumeroAgencia = 4334

	fmt.Println(contaNova.Depositar(100))
	fmt.Println(contaNova.Saldo())

	PagarBoleto(&contaNova, 45)
	fmt.Println(contaNova.Saldo())
}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
