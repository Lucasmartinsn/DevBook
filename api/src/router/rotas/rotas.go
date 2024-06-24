package rotas

import (
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/middlewares"
	"github.com/gorilla/mux"
)

type Rota struct {
	Uri                 string
	Metodo              string
	Funcao              func(http.ResponseWriter, *http.Request)
	ResquerAutntication bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUser
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {
		if rota.ResquerAutntication {
			r.HandleFunc(rota.Uri, 
				middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}
