package modelos

// Senha representa a requisição de troca de senha do usuário
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"Atual"`
}
