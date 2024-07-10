package rotas

import (
	"net/http"
	"webapp/src/controller"
)

// Rotas que manipular dados junto da api
var RotaPublicacaoCrud = []Rota{
	{
		URI:    "/publicacoes",
		Method: http.MethodPost,
		Funcao: controller.CriarPublicacao,
		Auth:   true,
	},
	{
		URI:    "/publicacoes/{id}/curtir",
		Method: http.MethodPost,
		Funcao: controller.CurtirPublicacao,
		Auth:   true,
	},
	{
		URI:    "/publicacoes/{id}/descurtir",
		Method: http.MethodPost,
		Funcao: controller.DesCurtirPublicacao,
		Auth:   true,
	},
	{
		URI:    "/publicacoes/{id}",
		Method: http.MethodPost,
		Funcao: controller.EditarPublicacao,
		Auth:   true,
	},
}
// Rotas que carregam paginas html
var RotaPublicacao = []Rota{
	{
		URI:    "/publicacao/{id}/editar",
		Method: http.MethodGet,
		Funcao: controller.CarregarPageEditarPublicacao,
		Auth:   true,
	},
}

