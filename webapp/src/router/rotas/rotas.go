package rotas

import (
	"net/http"

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
	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Method)
	}

	// Esta estrutura vai possibilitar que o Go consiga passar para o Navegador os Styles CSS
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
