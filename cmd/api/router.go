package main

func (app *application) routes() {
	app.router.GET("/index", app.indexHandler)
	app.router.GET("/dbtest", app.dbtestHandler)

}