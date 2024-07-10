package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/service"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// Essa funcao vai retorna a tela de login da aplicacao
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecultarTemplate(w, "login", nil)
}

// Essa funcao vai retorna a tela de cadastro da aplicacao
func CarregarPageCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecultarTemplate(w, "cadastroUser", nil)
}

// Essa funcao vai retorna a tela home da aplicacao
func CarregarPageHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%spublicacoes", os.Getenv("BASE_URL"))
	response, err := requisicoes.FazerRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	var publicacoes []models.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		respostas.Json(w, 422, respostas.ErrorApi{Error: err.Error()})
		return
	}

	cookie, _ := service.Ler(r)
	usuarioId, err := strconv.ParseUint(cookie["id"], 10, 64)
	if err != nil {
		respostas.Json(w, 422, respostas.ErrorApi{Error: err.Error()})
		return
	}
	utils.ExecultarTemplate(w, "home", struct {
		Publicacao []models.Publicacao
		Id         uint64
	}{
		Publicacao: publicacoes,
		Id:         usuarioId,
	})
}

// Essa funcao vai retorna a tela de cadastro da aplicacao
func CarregarPageEditarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	postId, err := strconv.ParseUint(parametro["id"], 10, 64)
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	response, err := requisicoes.FazerRequestWithAuth(r, http.MethodGet, fmt.Sprintf("%spublicacoes/%d", os.Getenv("BASE_URL"), postId), nil)
	if err != nil {
		respostas.Json(w, 500, respostas.ErrorApi{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}
	var publicacao models.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacao); err != nil {
		respostas.Json(w, 422, respostas.ErrorApi{Error: err.Error()})
		return
	}
	utils.ExecultarTemplate(w, "editarPost", publicacao)
}