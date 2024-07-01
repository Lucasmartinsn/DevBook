package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var RotasUsuario = []Rota{
	{
		URI:    "/criar-usuario",
		Method: http.MethodGet,
		Funcao: controller.CarregarPageCadastro,
		Auth:   false,
	},
}
