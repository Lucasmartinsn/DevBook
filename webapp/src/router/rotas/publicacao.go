package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var RotaPublicacao = []Rota{
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
}
