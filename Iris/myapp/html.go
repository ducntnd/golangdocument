package main

import "github.com/kataras/iris/v12"

func main() {

	app := iris.New()

	//config := iris.WithConfiguration(iris.Configuration {
	//	DisableStartupLog: true,
	//	EnableOptimizations: true,
	//	Charset: "UTF-8",
	//})

	config := iris.WithConfiguration(iris.YAML("./iris.yml"))

	app.RegisterView(iris.HTML("./html", ".html"))

	app.Use(middleWare)

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("fullname", "Nhat Duc")
		ctx.View("appgo.html")
		//ctx.ServeFile("html/appgo.html")
	})

	app.Get("/text", func(ctx iris.Context) {
		ctx.Text("Hello this is format text")
	})
	//
	//app.Get("/writes",func(ctx iris.Context) {
	//	ctx.WriteString("<h3>Hello this is Nhat Duc</h3>")
	//})
	//
	//app.Get("/json",func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{
	//		"title":"Mastermind",
	//	})
	//})

	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		lastname := ctx.URLParam("lastname") // shortcut for ctx.Request().URL.Query().Get("lastname")
		ctx.Writef("Hello %s %s", firstname, lastname)
	})

	none := app.None("/invisible/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s with method: %s", ctx.Params().Get("username"), ctx.Method())

		if from := ctx.Values().GetString("from"); from != "" {
			ctx.Writef("\nI see that you're coming from %s", from)
		}
	})

	app.Get("/execute", func(ctx iris.Context) {
		if !none.IsOnline() {
			ctx.Values().Set("from", "/execute with offline access")
			ctx.Exec("NONE", "/invisible/iris")
			return
		}
		ctx.Values().Set("from", "/execute")
		ctx.Exec("GET", "/invisible/iris")
	})
	app.Get("/change", func(ctx iris.Context) {

		if none.IsOnline() {
			none.Method = iris.MethodNone
		} else {
			none.Method = iris.MethodGet
		}
		app.RefreshRouter()
	})

	app.Run(iris.Addr(":8080"), config)
}

func middleWare(ctx iris.Context) {
	ctx.Application().Logger().Infof("Run before %s", ctx.Path())
	ctx.Next()
}
