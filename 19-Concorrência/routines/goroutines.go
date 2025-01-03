package main

import (
	"fmt"
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
	go escrever("Primeiro Texto")
	escrever("Segundo Texto")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
