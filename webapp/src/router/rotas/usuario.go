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
	{
		URI:    "/perfil/{id}",
		Method: http.MethodGet,
		Funcao: controller.CarregarPagePerfil,
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
	{
		URI:    "/usuario/{id}",
		Method: http.MethodPut,
		Funcao: controller.AtualizarDataUser,
		Auth:   false,
	},
	{
		URI:    "/usuario/{id}/atualizar-pass",
		Method: http.MethodPut,
		Funcao: controller.AtualizarPassUser,
		Auth:   false,
	},
}
