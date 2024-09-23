package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoKaiqui := ContaCorrente{titular: "Kaiqui", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	fmt.Println(contaDoKaiqui)

	contaDoRafael := ContaCorrente{"Rafael", 222, 111222, 200}

	fmt.Println(contaDoRafael)

	contaDaJu := ContaCorrente{}
	contaDaJu.titular = "Ju"
	contaDaJu.numeroAgencia = 555
	contaDaJu.numeroConta = 654321
	contaDaJu.saldo = 1050.56

	fmt.Println(contaDaJu)
}
