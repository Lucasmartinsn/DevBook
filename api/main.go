package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/config"
	"github.com/Lucasmartinsn/DevBook/api/src/router"
)

func main() {
	if err := config.Carregar(); err != nil {
		panic("Falha ao carregar variaveis de ambiente")
	}
	r := router.Generator()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
