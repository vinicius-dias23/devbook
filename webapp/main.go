package main

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplate()
	r := router.Gerar()
	fmt.Printf("Escutando na porta :%d\n", config.Porta)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)
}
