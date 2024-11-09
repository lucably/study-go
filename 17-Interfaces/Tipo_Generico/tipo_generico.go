package main

import "fmt"

//Um tipo de interface generica que aceita qualquer parametro.
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "abc",
		float32(100): true,
		"dfe":        1,
		true:         123,
	}

	fmt.Println(mapa)
}
