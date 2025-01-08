package enderecos

import "strings"

// TipoEndereco como será exportada, lembrar de colocar com letra maiuscula
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoMinuscula, " ")[0]

	enderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		//Coloca a primeira letra maiuscula
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"
}
