package controllers

import "net/http"

// CarregarTelaLogin renderiza a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de login"))
}
