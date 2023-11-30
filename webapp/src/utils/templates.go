package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func CarregarTemplate() {
	templates = template.Must(template.ParseGlob("view/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
