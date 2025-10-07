package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"html/template"
	"github.com/aikwen/webapp/static"
	"github.com/gin-gonic/gin"
)


type application struct {
	router *gin.Engine
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
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
