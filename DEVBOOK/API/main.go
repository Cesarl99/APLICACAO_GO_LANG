package main

import (
	"API/src/config"
	"API/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println("rodando API!!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
