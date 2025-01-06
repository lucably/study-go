package main

import "fmt"

func main() {
	/*
		Como estou chamando o canal na main, ele fica aguardando nesta mesma linha
		receber esse valor e mesmo que eu esteja passando o valor na linha 21
		isso nunca vai acontecer pois ele fica "travado" esperando o valor.
		Acontecendo o deadlock

		Agora passando uma capacidade, no caso capacidade 2 => canal := make(chan string, 2)
		isso quer dizer que o canal só irá bloquear quando chegar na capacidade máxima dele, no caso 2.
		Passando 2 valores, repare que ele não irá travar mas não ira printar os 2 valores pois irá passar
		direto e finalizar o programa depois da 1 execução. Caso passe mais um valor exp: "canal <- "Texto 3"" , irá acontecer o deadlock.
	*/
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Estudando Go!"

	mensagem := <-canal
	//Para printar o 2 valor, "Estudando Go!" vai precisar criar um novo canal, conforme o exp abaixo.
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
