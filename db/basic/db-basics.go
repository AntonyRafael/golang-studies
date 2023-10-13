package another

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func another() {
	stringConexao := "golang:golang@/devbooks?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		log.Fatal(err.Error())
	}
	// defer é executado no final da função
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Conexão está aberta!")

	// Executando comandos
	linhas, err := db.Query("select * from usuarios")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer linhas.Close()
	fmt.Println(linhas)
}
