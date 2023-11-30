package main

import (
	"fmt"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplate()
	r := router.Gerar()
	fmt.Println("Escutando na porta :7778")
	http.ListenAndServe(":7778", r)
}
