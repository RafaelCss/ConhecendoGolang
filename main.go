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
	fmt.Println(contaCorrenteRafael.Sacar(50.10), contaCorrenteRafael.saldo)
	fmt.Println(contaCorrenteRafael.Depositar(250))
	contaCorrenteRafael.ConsultarSaldo(contaCorrenteRafael.saldo)
	contaCorrenteRafael.VerificarSaldo()
}

/* Criando métodos de uso : */

func (conta * ContaCorrente) Sacar (valorDoSaque float64) string{
	var verificarSaldoPositivo = valorDoSaque > 0 && valorDoSaque <= conta.saldo
	if !verificarSaldoPositivo {
		return "Saldo insuficiente"
	}
	conta.saldo -= valorDoSaque
	return "Valor Sacado foi de ..."
}

func (conta * ContaCorrente) Depositar (valorDoDeposito float64) string{
	if valorDoDeposito > 0{
		conta.saldo += valorDoDeposito
		return "Valor Depositado foi de ..."
	}
	return "Valor não pode ser negativo"
}

func (conta * ContaCorrente) VerificarSaldo(){
	if conta.saldo > 0 {
		fmt.Println("Saldo positivo")
	}
		if conta.saldo <= 0 {
		fmt.Println("Sem Saldo")
	}

}



func (conta * ContaCorrente) ConsultarSaldo(saldo float64){
	fmt.Println("Saldo atual:",saldo)
}