package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Chamando a funcao escrever() nao precisara chamar a palavra reservada "go"
		A funcao escrever() que esta encapsulando a goroutines
	*/

	//Esse "canal" e do tipo que so recebe.
	canal := escrever("Olá Mundo!")

	//utilizando o generator escondemos a complexidade do codigo
	for i := 0; i < 10; i++ {
		fmt.Println("Printando sem o <-: ", canal)
		fmt.Println("Printando Normal: ", <-canal)
	}
}

// Vai retornar um canal de string
func escrever(texto string) <-chan string {
	canal := make(chan string)

	//goroutines rodando infinito
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Microsecond * 500)
		}
	}()

	return canal
}
