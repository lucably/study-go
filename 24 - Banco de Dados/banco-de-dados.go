package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// _ "github.com/go-sql-driver/mysql" => Import implicito, se vc tentasse colocar o modulo do mysql depois que salva-lo ele iria remove-lo pois n estaria sendo usado neste arquivo.
// O Pacote que realmente esta usando o github.com/go-sql-driver/mysql é o  "database/sql"

func main() {
	//"usuario:senha@/banco"
	urlConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		log.Fatal(erro)
	}
	//Fechando a conexão depois da conexão;
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	//Consulta banco
	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}

	//Fechando a conexão depois da consulta;
	defer linhas.Close()

	fmt.Println("Linhas: ", linhas)
}
