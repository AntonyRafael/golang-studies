package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa uma rota da aplicação web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar configura as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return router
}
