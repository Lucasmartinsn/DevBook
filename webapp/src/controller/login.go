package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"webapp/src/models"
	"webapp/src/respostas"
	"webapp/src/service"
)

// Essa função vai fazer a requisição de login
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	// Ele vai pegar o Bode da request e vai deixar os dados acessiveis para manipulação
	r.ParseForm()
	// Capturando os dados do Body e transformando em um array de Bytes
	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	response, err := http.Post(fmt.Sprintf("%slogin", os.Getenv("BASE_URL")), "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	var DadosAuth models.DadosAuth
	if err = json.NewDecoder(response.Body).Decode(&DadosAuth); err != nil {
		respostas.Json(w, http.StatusUnprocessableEntity, respostas.ErrorApi{Error: err.Error()})
		return
	}

	id := strconv.FormatUint(DadosAuth.IdUser, 10)
	if err = service.Salvar(w, id, DadosAuth.Token); err != nil {
		respostas.Json(w, http.StatusUnprocessableEntity, respostas.ErrorApi{Error: err.Error()})
		return
	}
	respostas.Json(w, response.StatusCode, nil)
}
