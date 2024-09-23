package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoKaiqui := ContaCorrente{}
	contaDoKaiqui.titular = "Kaiqui"
	contaDoKaiqui.saldo = 400

	fmt.Println("Sr.", contaDoKaiqui.titular, "seu saldo é:", contaDoKaiqui.saldo)

	valorDoSaque := 300.

	fmt.Println(contaDoKaiqui.sacar(valorDoSaque))

	valorDoDeposito := 500.

	fmt.Println(contaDoKaiqui.depositar(valorDoDeposito))
}

func (c *ContaCorrente) sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso! Saldo atual:", c.saldo
	} else {
		return "Saldo insuficiente! Saldo atual:", c.saldo
	}
}

func (c *ContaCorrente) depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso! Saldo atual:", c.saldo
	} else {
		return "Deposito Negado! Saldo atual:", c.saldo
	}
}
