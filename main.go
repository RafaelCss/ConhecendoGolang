package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroConta   int
	numeroAgencia int
	saldo         float64
}

func main() {
	contaCorrenteRafael := ContaCorrente{"Rafael", 599, 5998, 150.5}
	fmt.Println(contaCorrenteRafael)
}