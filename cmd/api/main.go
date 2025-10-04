package main

import (
	"html/template"
	"github.com/aikwen/webapp/static"
	"github.com/gin-gonic/gin"
)


type application struct {
	router *gin.Engine
}


func main() {
	app := &application{
		router: gin.Default(),
	}

	templ := template.Must(static.GetTemplates())
  	app.router.SetHTMLTemplate(templ)

	app.routes()
	app.router.Run(":8080")
}
