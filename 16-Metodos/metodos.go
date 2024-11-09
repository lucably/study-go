package main

import "fmt"

//Metodos pertecem a um Struct

type usuario struct {
	nome  string
	idade uint8
}

/*
	To dizendo que todo usuario tem o metodo salvar
	A variavel "u" => serve para buscar os dados do elemento q chamou a função
*/
func (u usuario) salvar() {
	fmt.Println("Salvando o usuario: ", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

/*
	Essa função adicionaria mais 1 na idade

	Se você for fazer uma comparação que não altere nenhum valor do struct
	não precisa passar como ponteiro
*/
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Lucas", 28}
	usuario2 := usuario{"Walter", 17}

	fmt.Println(usuario1)
	usuario1.salvar()

	retornaSeEMaiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(retornaSeEMaiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println("Usuario: ", usuario1)

	fmt.Println("Usuario2: ", usuario2)
}
