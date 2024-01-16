package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver do mysql
)

// Conectar vai conectar com o banco de dados e retornar o banco de dados e um erro
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBancoDeDados)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
