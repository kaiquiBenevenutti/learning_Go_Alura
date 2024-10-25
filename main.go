package main

import (
	"fmt"

	"github.com/learning_Go_Alura/contas"
)

func main() {
	contaDoRafael := contas.ContaCorrente{Titular: "Rafael", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 500}

	status := contaDoGustavo.Transferir(100, &contaDoRafael)

	fmt.Println(status)

	fmt.Println(contaDoGustavo)
	fmt.Println(contaDoRafael)
}
