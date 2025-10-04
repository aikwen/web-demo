package static

import (
	"embed"
	"html/template"
)

//go:embed *.html
var templateFS embed.FS

func GetTemplates() (*template.Template, error) {
	return template.ParseFS(templateFS, "*.html")
}