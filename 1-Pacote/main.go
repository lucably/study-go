package main

//Sempre que for usar algo importado utilizar o nome depois da ultima "/"
import (
	"fmt"
	auxiliar "modulo-dentro-do-1-pacote/pacote-auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devtestemail@gmail.com")
	fmt.Println(erro) //ele ira retornar "<nil>" que é nulo, ou seja, email válido.
}

/*
INFOS:

1) Caso queira gerar o arquivo build binario
digite o comando:
	go build
e para executar:
	./{arquivo_gerado}

2) instalar pacote: go get {caminho_url}
	exp: go get github.com/badoux/checkmail

obs: adiciona no arquivo go.mod

3) Remover todos os pacotes não utilizados:
	go mod tidy
*/
