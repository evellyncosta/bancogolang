package main

import (
	"fmt"

	"github.com/evellyncosta/bancogolang/clientes"
	"github.com/evellyncosta/bancogolang/contas"
)

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.123.123-12",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123,
		NumeroConta:   12345,
		Saldo:         100,
	}

	fmt.Println(contaDoBruno)

}
