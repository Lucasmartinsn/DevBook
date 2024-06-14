package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Porta = 0

func Carregar() (err error) {
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error ao crrega valores de usuario")
	}
	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 5000
	}
	return
}
