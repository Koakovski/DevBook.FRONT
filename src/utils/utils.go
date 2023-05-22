package util

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

func ExecTemplate(w http.ResponseWriter, templateName string, data any) {
	templates.ExecuteTemplate(w, templateName, data)
}
