package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome                string `json:"nome"`
	Raca                string `json:"raca"`
	Idade               uint   `json:"idade"`
	IgnorarAlgoDoStruct bool   `json:"-"`
}

func main() {

	c := cachorro{Nome: "Rex", Raca: "Vira Lata", Idade: 5}

	cachorroParaJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Em Bytes: ", cachorroParaJSON)
	fmt.Println("Em JSON: ", bytes.NewBuffer(cachorroParaJSON))

	cachorroEmJSON := `{
		"nome": "Rex",
		"raca": "Dálmata",
		"idade": 3
	}`

	var c2 cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Formato de Struct: ", c2)
}
