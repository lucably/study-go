package main

import "fmt"

func main() {
	numero := 10

	/*
		Explicando o IF INITI
		Criação da VARIAVEL outroNumero somente no escopo IF.
	*/

	if outroNumero := numero; numero > 0 {
		fmt.Println("Outro numero é maior que 0, Valor de Outro Numero: ", outroNumero)
	} else if numero < 0 {
		fmt.Println("Outro numero é menor que 0, Valor de Outro Numero: ", outroNumero)
	} else {
		fmt.Println("Outro numero é 0, Valor de Outro Numero: ", outroNumero)
	}

	/*
		Existem os tipos de FOR normal (igual nas outras linguagens)
		e tem tb este tipo.
	*/

	nomesAnimais := [3]string{"Cachorro", "Gato", "Cavalo"}

	for i, v := range nomesAnimais {
		fmt.Println("Indice: ", i, "Valores: ", v)
	}

	//Se quiser somente os Valores:
	for _, v := range nomesAnimais {
		fmt.Println("Valores: ", v)
	}
}
