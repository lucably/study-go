package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// BuscarUsuarios busca todos os usuários do DB.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		//  Scan => Pega cada linha e copia na ordem exata que veio.
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	//Transformando o slice em JSON.
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários em JSON"))
		return
	}
}

// BuscarUsuario busca um usuário do DB.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	//ParseUint => recebe 3 parametros: Variavel, baseNumero, TamanhoBit
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
	}

	//Abrir a conexao do banco aqui para ter a certeza que o usuario passou um parametro id.

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao DB."))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário."))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário."))
			return
		}
	}

	//Caso tente buscar um usuario que não exista ele popula como default para ambos o parametro, como ID é int, com o valor 0, Nome string, com o valor "", Email string, com o valor ""
	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}
}
