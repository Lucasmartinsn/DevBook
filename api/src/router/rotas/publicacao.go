package rotas

import (
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/controllers"
)

var rotasPublicacao = []Rota{
	{
		Uri:                 "/publicacoes",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarPublicacao,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/publicacoes",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoes,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/publicacoes/{id}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacao,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/publicacoes/{id}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarPublicacao,
		ResquerAutntication: true,
	},
	{
		Uri:                 "/publicacoes/{id}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarPublicacao,
		ResquerAutntication: true,
	},
}
