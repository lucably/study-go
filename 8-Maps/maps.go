package main

import "fmt"

func main() {

	// map[TIPO_CHAVE] TIPO_VALOR
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println("Valor: ", usuario["nome"])

	usuario2 := map[string]map[int]string{
		"nome": {
			1: "Joao",
			2: "Pedro",
		},
		"curso": {
			1: "Engenharia",
			2: "Matematica",
		},
	}

	fmt.Println("Usuario2: ", usuario2)
	fmt.Println("Valor: ", usuario2["nome"])
	fmt.Println("Valor final: ", usuario2["nome"][2])

	//Se quiser deletar uma chave tem o "delete" nativo do go.

	delete(usuario2, "nome")
	fmt.Println("Usuario2: ", usuario2)
	fmt.Println("Valor: ", usuario2["curso"])
	fmt.Println("Valor final: ", usuario2["curso"][2])
}
