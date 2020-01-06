package templateManager

import (
	"io"
	"html/template"
)

var templates *template.Template

func InitializeTemplateEngine() {
	templates = template.Must(template.ParseGlob("public/templates/*.html"))
}

func RenderHtmlWithParameters(w io.Writer, templatePath string, parameterStruct interface{}) {
	templates.ExecuteTemplate(w, templatePath, parameterStruct)
}