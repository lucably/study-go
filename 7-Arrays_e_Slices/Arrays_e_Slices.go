package main

import "fmt"

func main() {

	//Todos os dados do array tem que ser do msm tipo.
	var array [5]int
	array[2] = 3
	fmt.Println(array)

	array2 := [5]string{"Posição 2", "Posição 3"}
	fmt.Println(array2)

	/*
		[...] => serve para dizer que o tamanho do array é baseado na quantidade de elementos passados.
		Obs: Isso nao deixa o array dinamico.
	*/
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	/*
		Arrays não são muito utilizados por conta deste problema.
		O que utilizam mais é o Slice.
	*/

	/*
		A grande diferença do slice é que n dizemos o seu tamanho.
		Continua podendo ter somente um tipo de dado dentro dele.
	*/
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println("Slice antigo: ", slice)

	slice = append(slice, 16)
	fmt.Println("Slice novo: ", slice)

	/*
		Pegar uma certa quantidade do Slice.
		slice[1:5] => pega o elemento da posição 1 e vai percorrendo pegando até chegar
					na posição antes da final, no caso 5-1 = 4; Pegando os valores da
					posicao 1,2,3 e 4
	*/
	slice2 := slice[1:5]
	fmt.Println("Slice2 : ", slice2)

	/*
		make(TIPO, TAM, CAP) => aloca um espaço de memoria
	*/

	slice3 := make([]float32, 10, 15)
	fmt.Println("slice3: ", slice3)
	fmt.Println("Length: ", len(slice3))
	fmt.Println("Capacidade: ", cap(slice3))

	//Estourar o make.
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)
	slice3 = append(slice3, 3)
	slice3 = append(slice3, 4)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println("slice3: ", slice3)
	fmt.Println("Length: ", len(slice3))
	fmt.Println("Capacidade: ", cap(slice3))

	//Podemos tb fazer desse jeito, sem passar a CAP
	slice4 := make([]float32, 5)
	fmt.Println("slice4: ", slice4)
	fmt.Println("Length: ", len(slice4))
	fmt.Println("Capacidade: ", cap(slice4))
	slice4 = append(slice4, 1)
	fmt.Println("slice4: ", slice4)
	fmt.Println("Length: ", len(slice4))
	fmt.Println("Capacidade: ", cap(slice4))
}
