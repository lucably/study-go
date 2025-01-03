package main

import (
	"fmt"
	"time"
)

func main() {

	//chan => channels, canais
	//make => Utilizado para inicializar "slices, maps, channels"; make(TIPO, TAM, CAP) => aloca um espaço de memoria
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	/*
		//Recebendo um valor, esperando que o canal receba um valor.
		//Ele para aqui e espera até o canal ter um valor (no caso depois que passar da execução da função ESCREVER)
		mensagem := <-canal

		Se eu deixar sem esse "for" infinito, o canal só é executado 1 vez, pois ele espera receber algum retorno(valor)
		depois de ter recebido ele ja mostra e continua a execução, com o "for" obrigamos ele a esperar terminar de receber TODOS os valores.

		"canal" é utilizado igual ao "waitgroup"


		for {
			//"canal" tem uma segunda variavel que serve para checkar se ainda esta "aberto" o canal
			mensagem, aberto := <-canal

			if !aberto {
				//Se o "canal" n estiver aberto, eu saio do looping.
				break
			}

			fmt.Println(mensagem)
		}

		Uma outra maneira de fazer o percuso sem precisar fazer o comando "for" e o "<-canal" é assim:
		Dessa maneira n precisamos colocar o IF nem checar se ainda está aberto.
	*/
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//Mandando um valor para dentro do "canal"
		canal <- texto
		time.Sleep(time.Second)
	}
	//Quando terminar o looping, no caso de 5x, fecha(termina) o canal.
	close(canal)
}

/*
	SEMPRE que usar "chan" (canal) utiliza o "close" e a segunda variavel do retorno do carnal que chamamos de "aberto"

	Caso nao utilize, ou esqueça de utilizar, o programa quebra e retorna um "deadlock" e esses deadlock só é visto em PRODUÇÃO ele deixa passar na hora do build
	que causa quebra em produção.

	Mesmo o comando "go" serve para ir executando as outras linhas, o programa é parado pelo comando do canal quando espera receber o valor "<- canal".
*/
