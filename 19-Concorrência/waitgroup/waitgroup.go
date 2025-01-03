package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		Com a palavra reservada "go" faz com que o programa siga
		para a linha de baixo sem esperar que a linha atual termine
		sua execução.

		Removendo o "go" fara com que o programa trave executando somente
		a frase "Primeiro Texto"
	*/

	var waitGroup sync.WaitGroup

	//Criei 2 rotinas que estou dizendo para ele esperar terminar. Nao falo quem e sim quantas ele tem que esperar.
	waitGroup.Add(2)

	// "go" juntamente com Função anonima
	go func() {
		escrever("Primeiro Texto")

		//Diz que terminou a rotina
		waitGroup.Done()
	}()

	go func() {
		escrever("Segundo Texto")
		waitGroup.Done()
	}()

	//So sai daqui depois que tiver executado 2x waitGroup.Done()
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
