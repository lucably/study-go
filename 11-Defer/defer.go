package main

import "fmt"

func funcao1() {
	fmt.Println("executando a função 1")
}

func funcao2() {
	fmt.Println("executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	defer fmt.Println("Media calculada. Resultado: ", media)
	fmt.Println("Calculando a Média....")

	if media >= 6 {
		/*
			fmt.Println("Media calculada. Resultado: ", media)

			Poderia ter colocado aqui e repetir o codigo no trecho
			TRUE e para FALSE.
			Mas Podemos usar o DEFER, que fará printar somente quando tudo
			deste espoco for executado.
			Obs: Sera mostrando por ultimo antes do RETURN
		*/
		return true
	}
	/*
		fmt.Println("Media calculada. Resultado: ", media)

		Poderia ter colocado aqui e repetir o codigo tanto quando
		desse TRUE e para FALSE.
		Mas Podemos usar o DEFER, que fará printar somente quando tudo
		deste espoco for executado.
		Obs: Sera mostrando por ultimo antes do RETURN
	*/
	return false
}

// defer => Adia a execução desta função até o ultimo momento possivel.
// Muito utilizado para relações com o Banco de Dados.
func main() {
	defer funcao1()
	funcao2()
	fmt.Println(alunoAprovado(6, 7))
}
