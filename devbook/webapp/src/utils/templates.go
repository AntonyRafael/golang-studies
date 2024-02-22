package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates cria os templates a partir dos arquivos HTML
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate renderiza a pagina html com os dados
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
