package main

import "fmt"

func main() {
	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	/*
		Repare que somente a variavel1 alterou, pois a variavel2 recebeu a copia de seu valor
		e depois da linha 7 a variavel2 mantera o valor 10 até que seja alterada.
	*/
	fmt.Println(variavel1, variavel2)

	var variavel3 int
	//declaração do ponteiro. Por default o valor do ponteiro é null "<nil>"

	var ponteiro *int
	variavel3 = 100

	/*
		Se eu colocar exp => ponteiro = variavel3, nao dara certo.
		Para o ponteiro pegar o VALOR da variavel utilizamos o "&".
	*/
	ponteiro = &variavel3
	fmt.Println("Valor do ponteiro: ", ponteiro)
	/*
		O Ponteiro retornará este valor: "Valor do ponteiro:  0xc00000a120"
		No print nao estamos pegando o valor do ponteiro, e sim, onde ele salvou no
		endereço de memoria o valor 100.

		Para conseguir ver o VALOR que esta no ponteiro devemos utilizar a DESREFERENCIAÇÃO
		fmt.Println("Valor do ponteiro: ", *ponteiro)
	*/
	fmt.Println("Valor do ponteiro: ", *ponteiro)

}
