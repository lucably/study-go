package main

import "fmt"

var n int

func main() {
	fmt.Println("Função Main sendo executada")
	fmt.Println("Valor de N: ", n)
}

//Pode ter uma por arquivo.
//Utilizada para fazer setup ou utilizar para criação de variavel globais, iniciar variaveis
func init() {
	fmt.Println("Função Init sendo executada")
	n = 10
}
