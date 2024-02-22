package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaLogin renderiza a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
