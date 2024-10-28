package auxiliar

import "fmt"

/*Função com letra
minuscula: é visivel somente no pacote onde ela se encontra.
Maiuscula: Pode ser exportada para outros pacotes.
*/
func Escrever() {
	fmt.Println("Escrevendo a frase de um arquivo auxiliar do pacote AUXILIAR.")
	//Aqui vc pode chamar com letra minuscula pq o arquivo axuiliar2.go esta dentro do mesmo pacote onde se encontra o axuiliar.go
	escrever2()
}
