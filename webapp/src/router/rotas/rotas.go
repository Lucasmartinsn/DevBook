package rotas

import (
	"net/http"
	"webapp/src/middleweres"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da Aplicação
type Rota struct {
	URI    string
	Method string
	Funcao func(http.ResponseWriter, *http.Request)
	Auth   bool
}

// Essa função vai recuperar todas as rotas que tenhao sido criadas, e colocalas em um slice e retornaras
func Configurar(router *mux.Router) *mux.Router {
	rotas := RotaLogin
	rotas = append(rotas, RotasUsuario...)
	rotas = append(rotas, RotasUsuarioCrud...)
	rotas = append(rotas, RotaPublicacao...)
	for _, rota := range rotas {
		if rota.Auth {
			router.HandleFunc(rota.URI, middleweres.Logger(middleweres.Autenticacao(rota.Funcao))).Methods(rota.Method)
		} else {
			router.HandleFunc(rota.URI, middleweres.Logger(rota.Funcao)).Methods(rota.Method)
		}
	}

	// Esta estrutura vai possibilitar que o Go consiga passar para o Navegador os Styles CSS
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
