package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger é um middleware que loga as requisições recebidas
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requisição recebida: %s %s %s", r.Method, r.URL.Path, r.Host)
		next(w, r)
	})
}

// Autenticar verifica a existência de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		proximaFuncao(w, r)
	}
}
