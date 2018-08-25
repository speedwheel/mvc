package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/speedwheel/mvc/datasource"
	"github.com/speedwheel/mvc/repositories"
	"github.com/speedwheel/mvc/web/controllers"
	"github.com/speedwheel/mvc/web/middleware"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	db, err := datasource.NewDatabase(datasource.Aerospike)
	if err != nil {
		app.Logger().Fatalf("error while loading the database: %v", err)
		return
	}
	defer db.Close()
	repositories.NewPropertyRepository(db)

	index := mvc.New(app.Party("{resource:path}"))
	index.Handle(new(controllers.IndexController))
	api := mvc.New(app.Party("api."))
	api.Router.Use(middleware.BasicAuth)
	api.Handle(new(controllers.ApiController))
	api.Party("/v1").Handle(new(controllers.ApiV1Controller))

	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("mvc.com:80"),
		// Disables the updater.
		iris.WithoutVersionChecker,
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
