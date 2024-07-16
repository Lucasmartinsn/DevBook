package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"os"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/service"

	"github.com/gorilla/mux"
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
	respostas.Json(w, 200, "success")
}

// Vai chamar a API para atualizar o Cadastro do usuario
func AtualizarDataUser(w http.ResponseWriter, r *http.Request) {
	// Ele vai pegar o Bode da request e vai deixar os dados acessiveis para manipulação
	r.ParseForm()
	// Capturando os dados do Body e transformando em um array de Bytes
	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	parametro := mux.Vars(r)
	userId, err := strconv.ParseUint(parametro["id"], 10, 64)
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}

	response, err := requisicoes.FazerRequestWithAuth(r, http.MethodPut, fmt.Sprintf("%susuario/%d", os.Getenv("BASE_URL"), userId), bytes.NewBuffer(usuario))
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	respostas.Json(w, 200, "success")
}

// Vai chamar a API para atualizar a senha do usuario
func AtualizarPassUser(w http.ResponseWriter, r *http.Request) {
	// Ele vai pegar o Bode da request e vai deixar os dados acessiveis para manipulação
	r.ParseForm()
	// Capturando os dados do Body e transformando em um array de Bytes
	usuario, err := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	parametro := mux.Vars(r)
	userId, err := strconv.ParseUint(parametro["id"], 10, 64)
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}

	response, err := requisicoes.FazerRequestWithAuth(r, http.MethodPut, fmt.Sprintf("%susuario/%d/atualizar-pass", os.Getenv("BASE_URL"), userId), bytes.NewBuffer(usuario))
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	respostas.Json(w, 200, "success")
}

// Essa função vai apagar o Cookie
func DeletarCookie(w http.ResponseWriter, r *http.Request) {
	service.Delete(w)
	http.Redirect(w, r, "/", http.StatusFound)
}
