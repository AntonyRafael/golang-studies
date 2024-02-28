package modelos

// DadosAutenticacao armazena os dados de autenticação do usuário
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
