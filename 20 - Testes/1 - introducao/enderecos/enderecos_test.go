package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//Para criar um test em Go precisa criar o arquivo deste formato =: {nomeArquivo_test}.go

// TESTE DE UNIDADE

// Test{nomeFuncao} Comando para rodar => go test
func TestTipoEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"avenida teste", "Avenida"},
		{"Estrada Bandeirantes", "Estrada"},
		{"Rodovia Imigrantes", "Rodovia"},
		{"Praça do Boticario", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperando %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
