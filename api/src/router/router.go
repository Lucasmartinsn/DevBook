package router

import (
	"github.com/Lucasmartinsn/DevBook/api/src/router/rotas"
	"github.com/gorilla/mux"
)

func Generator() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
