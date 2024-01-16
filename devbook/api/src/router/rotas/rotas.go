package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa uma rota da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar vai configurar as rotas da API
func Configurar(r *mux.Router) *mux.Router {
	for _, rota := range rotasUsuarios {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
