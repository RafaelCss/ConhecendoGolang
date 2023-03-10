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
	fmt.Println(contaCorrenteRafael.Sacar(50.10))
	fmt.Println(contaCorrenteRafael.saldo)
	fmt.Print(contaCorrenteRafael.Depositar(250))
	fmt.Println(contaCorrenteRafael.saldo)
}



func (conta * ContaCorrente) Sacar (valorDoSaque float64) string{
	var verificarSaldoPositivo = valorDoSaque > 0 && valorDoSaque <= conta.saldo
	if !verificarSaldoPositivo {
		return "Saldo insuficiente"
	}
	conta.saldo -= valorDoSaque
	return "Valor Sacado foi de ..."
}

func (conta * ContaCorrente) Depositar (valorDoDeposito float64) string{
	conta.saldo += valorDoDeposito
	return "Valor Depositado foi de ..."
}