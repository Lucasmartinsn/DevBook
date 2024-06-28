package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Vai Expor as rotas disponiveis na aplicação
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
