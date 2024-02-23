package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// CriarUsuario chama a API para criar um usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	response, erro := http.Post("http://localhost:5005/usuarios", "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		log.Fatal(erro)
	}

	defer response.Body.Close()
}
