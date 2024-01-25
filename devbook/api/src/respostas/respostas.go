package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON vai enviar uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, status int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro vai enviar uma resposta de erro em JSON para a requisição
func Erro(w http.ResponseWriter, status int, erro error) {
	JSON(w, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
