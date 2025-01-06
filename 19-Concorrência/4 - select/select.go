package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	/*
		Desse jeito conforme o codigo abaixo:

		mensagemCanal1 := <-canal1
		fmt.Println(mensagemCanal1)

		mensagemCanal2 := <-canal2
		fmt.Println(mensagemCanal2)

		Esta tendo uma saida:
		---------------------
		canal1
		canal2
		canal1
		canal2
		---------------------
		Ao inves de ser
		canal1
		canal1
		canal1
		canal1
		canal2
		canal1
		canal1
		canal1
		canal1
		canal2

		Pois como canal1 é em Milisegundos e o canal2 em segundos era para acontecer conforme o primeiro exemplo.
		Nao esta acontecendo da maneira correta pois quando ele recebe o valor do canal1 ele espera
		receber o canal2 para continuar a operacao, por isso, ele nao printa 4x canal1 e depois canal2

	*/

	//Como correcao usamos "SELECT" que funciona como se fosse um switch case no qual ele so faz a acao quando atender os parametros.
	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
