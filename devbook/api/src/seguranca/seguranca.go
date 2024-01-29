package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash gera um hash a partir de uma string
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha com seu hash
func VerificarSenha(senhaHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}
