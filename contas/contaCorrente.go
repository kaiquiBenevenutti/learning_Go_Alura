﻿package contas

import (
	"Primeiro_Projeto/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque concluido com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDeposito := valorDoDeposito > 0
	if podeDeposito {
		c.saldo += valorDoDeposito
		return "Depositado com sucesso : ", valorDoDeposito
	} else {
		return "Deposito insuficiente invalido : ", valorDoDeposito
	}
}

func (c *ContaCorrente) Saldo() (string, float64) {
	return "Saldo da conta :", c.saldo
}
