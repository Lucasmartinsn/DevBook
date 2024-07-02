package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"webapp/src/router"
	"webapp/src/utils"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error ao crrega valores default")
	}
}

func main() {
	utils.CarregarTempletes()
	r := router.Gerar()

	fmt.Println("rodando app web")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT_API")), r))
}
