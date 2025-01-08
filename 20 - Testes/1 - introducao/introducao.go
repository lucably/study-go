package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("avenida paulista")
	fmt.Println(tipoEndereco)
	tipoEndereco = enderecos.TipoEndereco("praça teste 22")
	fmt.Println(tipoEndereco)
}
