package rotas

import (
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/router/controllers"
)

var rotasUser = []Rota{
	{
		Uri:                 "/usuario",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarUser,
		ResquerAutntication: false,
	},
	{
		Uri:                 "/usuario",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscaUsers,
		ResquerAutntication: false,
	},
	{
		Uri:                 "/usuario/{id}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscaUser,
		ResquerAutntication: false,
	},
	{
		Uri:                 "/usuario",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUser,
		ResquerAutntication: false,
	},
	{
		Uri:                 "/usuario/{id}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletaUser,
		ResquerAutntication: false,
	},
}
