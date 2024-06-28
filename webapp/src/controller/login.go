package controller

import (
	"net/http"
	"webapp/src/utils"
)

// Essa funcao vai retorna a tela de login da aplicacao
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("tela login"))
	utils.ExecultarTemplate(w, "login", nil)
}
