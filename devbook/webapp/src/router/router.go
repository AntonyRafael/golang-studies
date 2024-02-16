package router

import "github.com/gorilla/mux"

// Gerar retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return r
}
