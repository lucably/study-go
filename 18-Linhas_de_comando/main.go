package main

import (
	"linha-de-compando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	//Metodo Run retorna um erro, entao seria bom ter um tratamento
	aplicacao.Run(os.Args)

	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}

}
