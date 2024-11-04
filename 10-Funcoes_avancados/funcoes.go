package main

import "fmt"

/*
	Poderia cadastrar desta maneira tb:

	func calculosMatematicos(n1, n2 int) (int, int) {
		soma := n1 + n2
		sub := n1 - n2
		return soma, sub
	}
*/

func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return soma, sub
}

/*
	Funcoes Variaticas => Podem receber vários parametros

	exp: func exemplo(VARIAVEL ...TIPO)
	obs: Só pode ter um parametro variaticas por função,
	e ele terá que ser obrigatoriamente o utltimo parametro.
*/
func calculosMatematicosInfinitosNumeros(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, value := range numeros {
		total += value
	}

	return total
}

/*
	Funções Recursivas
*/

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)

	totalDaSoma := calculosMatematicosInfinitosNumeros(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total da soma: ", totalDaSoma)

	/*
		Funcoes Anonimas => Podem receber vários parametros

		func () {

		}()

		Esses "()" no final sinaliza para o GO que quando ela acaba de declarar ja executa.
	*/

	func() {
		fmt.Println("Função Anonimas, declarou ja executou")
	}()

	func(nome string) {
		fmt.Println("Função Anonimas, declarou ja executou MAS com parametro NOME: ", nome)
	}("Lucas")

	retorno := func(nome string) string {
		return fmt.Sprintf("Recebeu o nome: %s", nome)
	}("Lucas")

	fmt.Println(retorno)

	posicao := uint(10)
	fmt.Println(fibonacci(posicao))

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
