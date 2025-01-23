package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint   `"json:id"`
	Nome  string `"json:nome"`
	Email string `"json:email"`
}

// CriarUsuario insere um usuário no DB.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT => Ajuda contra SQL INJECTION
	// Criando o comando de inserção sem passar os valores para ele "?"
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	//Substituindo os valores da inserção
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	//STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}
