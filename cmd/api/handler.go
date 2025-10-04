package main

import "github.com/gin-gonic/gin"

func (app *application) indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}