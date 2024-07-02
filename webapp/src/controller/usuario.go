package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"webapp/src/respostas"
)

// Base URL
var url = os.Getenv("BASE_URL")

// Vai chamar a API para a criaçao de um novo Cadastro de usuario
func CadastroOfUser(w http.ResponseWriter, r *http.Request) {
	// Ele vai pegar o Bode da request e vai deixar os dados acessiveis para manipulação
	r.ParseForm()
	// Capturando os dados do Body e transformando em um array de Bytes
	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
	}
	response, err := http.Post(fmt.Sprintf("%susuario", url), "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	respostas.Json(w, response.StatusCode, nil)
}
