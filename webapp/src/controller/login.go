package controller

import "net/http"

// Essa funcao vai retorna a tela de login da aplicacao
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tela login"))
}
