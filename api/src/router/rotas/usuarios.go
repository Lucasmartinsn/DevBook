package rotas

import "net/http"

var rotasUser = []Rota{
	{
		Uri:    "/usuario",
		Metodo: http.MethodPost,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		ResquerAutntication: false,
	},
	{
		Uri:    "/usuario",
		Metodo: http.MethodGet,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		ResquerAutntication: false,
	},
	{
		Uri:    "/usuario/{id}",
		Metodo: http.MethodGet,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		ResquerAutntication: false,
	},
	{
		Uri:    "/usuario",
		Metodo: http.MethodPut,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		ResquerAutntication: false,
	},
	{
		Uri:    "/usuario/{id}",
		Metodo: http.MethodDelete,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		ResquerAutntication: false,
	},
}
