package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução!")
}

func recuperarExecucaoComRecover() {
	//O valor do "recover()" aqui utilizado a variavel "r", não tiver entrado em "panic" ele retornara "nil" ou seja, ignorando oq esta dentro do IF
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	//Se comentar a função "recuperarExecucaoComRecover()" dara erro pois "panic" necessita do "recover"
	defer recuperarExecucaoComRecover()
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

/*
	panic => Interrompe o fluxo normal do programa, parando a execução (matando o programa).
	Obs: Antes do programa encerrar a execução ele chama todas as funções que tem "defer".
*/
//recover => Serve para tratar caso o programa entre em no estado "panic"

//defer => Printa de baixo para cima
func main() {
	fmt.Println(alunoAprovado(6, 6))

	//ignorando o panic e recover
	fmt.Println("---------------------------")
	fmt.Println(alunoAprovado(6, 7))
}
