package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Println("A área da forma é %0.2f", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main() {
	/*
		Chamar essa funcao "escreverArea(r)" sem criar os metodos
		antes da erro pois ela necessita que seja implementada o metodo area da
		interface forma

		"retangulo does not implement forma (missing method area)"

		Somente ira conseguir realizar a func "escreverArea" quando foi digitado
		os metodos da linha 22 e 26.

		r := retangulo{10, 15}
		escreverArea(r);
	*/

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
