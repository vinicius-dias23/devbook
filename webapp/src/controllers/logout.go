package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Excluir(w)
	http.Redirect(w, r, "/login", 302)
}
