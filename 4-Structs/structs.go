package main

import "fmt"

type usuario struct {
	nome      string
	idade     uint8
	estudando bool
	endereco  endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	/*
		printa todos os valores do obj no modo default
		println => (" ",0,false);
	*/
	fmt.Println("Sem valor: ", u)

	u.nome = "Lucas"
	u.idade = 28
	u.estudando = true
	fmt.Println("Com valor: ", u)

	//Outra maneira de criar, porém, desta maneira é obrigatório passar todos os campos
	usuario2 := usuario{"Nilton", 60, false, endereco{"Rio de Janeiro", 5}}
	fmt.Println("Usuario2: ", usuario2)

	//Outra maneira de criar, porém, desta maneira é obrigatório passar todos os campos
	var enderecoComplete endereco = endereco{"São Paulo", 10}
	usuario22 := usuario{"Nilton 22", 60, false, enderecoComplete}
	fmt.Println("usuario22: ", usuario22)

	//Quiser passar somente um dos parametros, Obs: O resto vem como default:
	usuario3 := usuario{nome: "Ana"}
	fmt.Println("usuario3: ", usuario3)

}
