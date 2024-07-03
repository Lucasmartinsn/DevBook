package rotas

import (
	"net/http"
	"webapp/src/controller"
)

// Rotas de Renderização das paginas Web
var RotasUsuario = []Rota{
	{
		URI:    "/criar-usuario",
		Method: http.MethodGet,
		Funcao: controller.CarregarPageCadastro,
		Auth:   false,
	},
	{
		URI:    "/home",
		Method: http.MethodGet,
		Funcao: controller.CarregarPageHome,
		Auth:   true,
	},
}

// Rotas de consumo e manipulação dos dados
var RotasUsuarioCrud = []Rota{
	{
		URI:    "/usuario",
		Method: http.MethodPost,
		Funcao: controller.CadastroOfUser,
		Auth:   false,
	},
}
