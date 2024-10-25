package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoRafael := ContaCorrente{titular: "Rafael", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 500}

	status := contaDoGustavo.transferir(100, &contaDoRafael)

	fmt.Println(status)

	fmt.Println(contaDoGustavo)
	fmt.Println(contaDoRafael)
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

func (c *ContaCorrente) transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
