package controller

import (
	"fmt"
	"net/http"
	"os"
	"webapp/src/requisicoes"
	"webapp/src/utils"
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
	url := fmt.Sprintf("%spublicacao", os.Getenv("BASE_URL"))
	response, err := requisicoes.FazerRequestWithAuth(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, err)
	utils.ExecultarTemplate(w, "home", nil)
}
