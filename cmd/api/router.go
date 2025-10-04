package main

func (app *application) routes() {
	app.router.GET("/index", app.indexHandler)
}