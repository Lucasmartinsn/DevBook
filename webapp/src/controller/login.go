package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/respostas"
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
	response, err := http.Post(fmt.Sprintf("%slogin", url), "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	
	if token, err := io.ReadAll(response.Body); err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}else {
		respostas.Json(w, response.StatusCode, string(token))
	}
}
