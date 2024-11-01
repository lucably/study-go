package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

/*
Para realizar a "herança" basta colocar o nome do struct como se fosse variavel (sem o tipo)

exp:

	type estudante struct {
		pessoa
		curso string
		faculdade string
	}

Deste jeito, poderá acessar o campo "nome" de PESSOA somente com => estudante.nome

Já fazendo do jeito VARIAVEL TIPO.

exp:

	type estudante struct {
		pessoa pessoa
		curso string
		faculdade string
	}

ficaria assim a chamada => estudante.pessoa.nome
*/

func main() {
	p1 := pessoa{"Joao", "Pedro", 20, 178}
	fmt.Println("p1: ", p1)

	e1 := estudante{p1, "Engenharia", "IF"}
	fmt.Println("Pegando somente o nome da p1: ", e1.nome)
}
