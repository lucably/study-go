package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

/*
calculosMatematicos(n1, n2 int8) => tanto n1 como n2 sao do tipo int8
quando nao se declara o ultimo parametro passa para todos os outros.

Funcoes com mais de 1 retorno => Na funcao abaixo teremos retorno de INT8 e outro retorno como BOOL
*/
func calculosMatematicos(n1, n2 int8) (int8, string) {
	soma := n1 + n2
	//Nao existe operador ternário no Go
	var parOuImpar string

	if soma%2 == 0 {
		parOuImpar = "par"
	} else {
		parOuImpar = "impar"
	}

	return soma, parOuImpar
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//Pode ser chamada tb deste modo -> f := func(texto string);;;
	var f = func(texto string) {
		fmt.Println(texto)
	}

	f("texto passado por parametro")

	resultadoValores, resultadoParOuImpar := calculosMatematicos(10, 10)
	resultadoValores2, resultadoParOuImpar2 := calculosMatematicos(10, 5)

	fmt.Println("resultadoValores: ", resultadoValores, "resultadoParOuImpar: ", resultadoParOuImpar)
	fmt.Println("resultadoValores2: ", resultadoValores2, "resultadoParOuImpar2: ", resultadoParOuImpar2)

	//Como a funcao necessita de 2 parametros (e é obrigado a passa-los, caso nao for usa-la utilize o "_");
	resultadoValoresSem2Parametro, _ := calculosMatematicos(10, 10)
	fmt.Println("resultadoValoresSem2Parametro: ", resultadoValoresSem2Parametro)
}
