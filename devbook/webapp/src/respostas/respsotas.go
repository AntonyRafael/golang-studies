package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa um erro que a API retorna
type Erro struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(dados)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// TratarStatusCodeDeErro trata os erros de status code
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erroAPI Erro
	json.NewDecoder(r.Body).Decode(&erroAPI)
	JSON(w, r.StatusCode, erroAPI)
}
