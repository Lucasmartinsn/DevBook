package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var RotaPublicacao = []Rota{
	{
		URI:    "/publicacoes",
		Method: http.MethodPost,
		Funcao: controller.CarregarPagePublicacao,
		Auth:   true,
	},
}
