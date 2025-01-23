package main

import (
	"crud/servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Contem todas as rotas que faz ligação com o DB.
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5000", router))
}