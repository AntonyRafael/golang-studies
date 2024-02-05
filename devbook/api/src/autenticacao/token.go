package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken cria um token com base no id do usuário
func CriarToken(id uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString(config.SecretKey)
}
