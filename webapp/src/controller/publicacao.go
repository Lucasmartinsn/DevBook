package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// Vai criar uma publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	// Ele vai pegar o Bode da request e vai deixar os dados acessiveis para manipulação
	r.ParseForm()
	// Capturando os dados do Body e transformando em um array de Bytes
	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	response, err := requisicoes.FazerRequestWithAuth(r, http.MethodPost, fmt.Sprintf("%spublicacoes", os.Getenv("BASE_URL")), bytes.NewBuffer(publicacao))
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	respostas.Json(w, response.StatusCode, nil)
}

// Vai curtit uma publicação
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	// Ele vai pegar o Bode da request e vai deixar os dados acessiveis para manipulação
	parametro := mux.Vars(r)
	postId, err := strconv.ParseUint(parametro["id"],10, 64)
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}

	response, err := requisicoes.FazerRequestWithAuth(r, http.MethodPost, fmt.Sprintf("%spublicacoes/%d/like", os.Getenv("BASE_URL"), postId), nil)
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	respostas.Json(w, response.StatusCode, nil)
}