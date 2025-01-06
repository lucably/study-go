package main

import "fmt"

func main() {
	//channels com buffer de 45 posicoes
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//Executando apenas 1 funcao demoraria muito tempo para fazer, chamando ela 4x a execucao comeca a ficar mais rapido.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

/*
Declaracao da funcao:

tarefas <- chan int => Quer dizer que tarefas é o que recebe dados
resultados chan<- int => Quer dizer que resultado é o que envia os dados
*/
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
