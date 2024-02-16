package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando webapp")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8080", r))
}
