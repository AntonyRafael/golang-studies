package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbooks?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
