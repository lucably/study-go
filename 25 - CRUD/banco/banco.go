package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com o MySQL
)

// Conectar abre a conexão com o banco de dados que retorna o DB e um erro.
func Conectar() (*sql.DB, error) {
	urlConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		return nil, erro //Retornando nil em relação ao DB, já que deu erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil //Retornado o DB e retornando nil no erro
}
