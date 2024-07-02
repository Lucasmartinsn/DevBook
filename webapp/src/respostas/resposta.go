package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct de resposta para tratar erros
type ErrorApi struct {
	Error string `json:"erro"`
}

// Retorna uma resposta em formato json para a request
func Json(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Vai tratar os Erros retornado no corpo(BODY) da requesição que foi retornada pela API
func TratarRespostaErro(w http.ResponseWriter, r *http.Response) {
	var erro ErrorApi
	json.NewDecoder(r.Body).Decode(&erro)
	Json(w, r.StatusCode, erro)
}
