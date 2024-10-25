package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso! Saldo atual:", c.Saldo
	} else {
		return "Saldo insuficiente! Saldo atual:", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso! Saldo atual:", c.Saldo
	} else {
		return "Deposito Negado! Saldo atual:", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
