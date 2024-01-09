package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotaHome = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaHome,
	RequerAutenticacao: true,
}
