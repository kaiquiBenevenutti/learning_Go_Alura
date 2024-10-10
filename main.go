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
	contaDoKaiqui.saldo = 600

	contaMarcos := ContaCorrente{}
	contaMarcos.titular = "Marcos"
	contaMarcos.saldo = 500

	fmt.Println("Sr.", contaDoKaiqui.titular, "seu saldo é:", contaDoKaiqui.saldo)
	fmt.Println("Sr.", contaMarcos.titular, "seu saldo é:", contaMarcos.saldo)

	status := contaDoKaiqui.transferir(300, &contaMarcos)

	if status {
		fmt.Println("Transferencia realizada com sucesso!")
	} else {
		fmt.Println("Transferencia negada!")
	}

	fmt.Println("Sr.", contaDoKaiqui.titular, "seu saldo é:", contaDoKaiqui.saldo)
	fmt.Println("Sr.", contaMarcos.titular, "seu saldo é:", contaMarcos.saldo)
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
	podeTransferir := valorDaTransferencia <= c.saldo && c != contaDestino

	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.saldo += valorDaTransferencia
		return true
	} else {
		return false
	}
}
