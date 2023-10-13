package main

import (
	"db/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CriarUsuario).Methods((http.MethodPost))
	router.HandleFunc("/usuarios", server.BuscarUsuarios).Methods((http.MethodGet))
	router.HandleFunc("/usuario/{id}", server.BuscarUsuario).Methods((http.MethodGet))
	router.HandleFunc("/usuario/{id}", server.AtualizarUsuario).Methods((http.MethodPut))
	router.HandleFunc("/usuario/{id}", server.DeletarUsuario).Methods((http.MethodDelete))

	fmt.Println("Escutando na porta 5005")

	log.Fatal(http.ListenAndServe(":5005", router))
}
