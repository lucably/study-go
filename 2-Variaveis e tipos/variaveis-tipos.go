package main

import (
	"errors"
	"fmt"
)

func main() {
	//Variavel tipada.
	var variavelStringTipada string = "Variável string"
	fmt.Println(variavelStringTipada)

	//Variavel oculto;
	variavelOculta := "Variável2 string "
	fmt.Println(variavelOculta)

	//outros modos de declarações
	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	//Trocar valores de variaveis
	x, y := 5, 10
	fmt.Println("X Anterior: ", x, "Y Anterior: ", y)
	x, y = y, x
	fmt.Println("X Depois: ", x, "Y Depois: ", y)

	var textoEmBranco string
	fmt.Println("textoEmBranco: ", textoEmBranco)

	var numeroEmBranco int
	fmt.Println("numeroEmBranco: ", numeroEmBranco)

	var boolEmBranco bool
	fmt.Println("boolEmBranco: ", boolEmBranco)

	//Erro em go é um tipo => Que retorna <nil>
	var errorEmBranco error
	fmt.Println("errorEmBranco: ", errorEmBranco)

	//errors => é um pacote no Go
	var error error = errors.New("Erro interno")
	fmt.Println("error: ", error)
}
