package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBancoDeDados vai conter a string de conexão com o banco de dados (mysql)
	StringConexaoBancoDeDados = ""
	// Porta vai conter a porta que o servidor vai escutar
	Porta = 0
	// SecretKey vai conter a chave secreta para criptografar os tokens
	SecretKey = []byte("")
)

// Carregar vai inicializar as configurações do projeto
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente")
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 5005
	}

	StringConexaoBancoDeDados = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
