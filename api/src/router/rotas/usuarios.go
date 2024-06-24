package rotas

import (
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/controllers"
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
		ResquerAutntication: true,
	},
	{
		Uri:                 "/usuario/{id}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscaUser,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/usuario/{id}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUser,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/usuario/{id}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletaUser,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/usuario/{id}/seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.SeguirUser,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/usuario/{id}/parar-de-seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.PararDeSeguir,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/usuario/{id}/seguidores",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarSeguidores,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/usuario/{id}/seguindo",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarSeguindo,
		ResquerAutntication: true,
	},
}
