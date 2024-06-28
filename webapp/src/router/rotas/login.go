package rotas

import (
	"net/http"
	"webapp/src/controller"
)

// Definindo as rotas de Login
var RotaLogin = []Rota{
	{
		URI:    "/",
		Method: http.MethodGet,
		Funcao: controller.CarregarTelaLogin,
		Auth:   false,
	},
	{
		URI:    "/login",
		Method: http.MethodGet,
		Funcao: controller.CarregarTelaLogin,
		Auth:   false,
	},
}
