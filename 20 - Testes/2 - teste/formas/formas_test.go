package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//t.Run() => Dividir o teste em 2 etapas.
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaEsperada {
			//diferente do Errorf o Fatalf quando encontra um erro ele ja para a execução.
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaEsperada {
			//diferente do Errorf o Fatalf quando encontra um erro ele ja para a execução.
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}

//Digitando go test -v retornará:
/*
=== RUN   TestArea
=== RUN   TestArea/Retângulo
=== RUN   TestArea/Circulo
--- PASS: TestArea (0.00s)
    --- PASS: TestArea/Retângulo (0.00s)
    --- PASS: TestArea/Circulo (0.00s)
PASS
ok      formas-testes/formas    0.212s

*/
