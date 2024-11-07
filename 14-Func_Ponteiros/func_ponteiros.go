package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	//Jogando um valor novo dentro do endereço de memoria
	*numero = *numero * -1
}

//Ponteiro seria mais utilizado para mudar uma variavel no programa inteiro.
func main() {
	numero := 20
	numeroIntertido := inverterSinal(numero)
	fmt.Println(numeroIntertido)
	//Repare que numero continua "20", pois o numeroIntertido recebe apenas a copia do numero
	fmt.Println(numeroIntertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("Novo numero com valor trocado na mesma variavel: ", novoNumero)
}
