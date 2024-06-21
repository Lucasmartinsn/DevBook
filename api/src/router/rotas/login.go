package rotas

import (
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/controllers"
)

var rotaLogin = Rota{
	Uri:                 "/login",
	Metodo:              http.MethodPost,
	Funcao:              controllers.Login,
	ResquerAutntication: false,
}
