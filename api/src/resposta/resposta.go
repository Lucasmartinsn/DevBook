package resposta

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Erro(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
