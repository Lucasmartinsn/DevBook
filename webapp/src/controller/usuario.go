package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"os"
	"webapp/src/models"
	"webapp/src/respostas"
)

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
		return
	}
	response, err := http.Post(fmt.Sprintf("%susuario", os.Getenv("BASE_URL")), "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	var dadosAutenticacao models.DatosAuth
	if err = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}

	respostas.Json(w, 200, nil)
}
