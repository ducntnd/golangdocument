package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./",".html"))
	// define a function
	h := func(ctx iris.Context) {
		ctx.View("index.html")
	}

	home := app.Get("/", h)
	home.Name = "home"

	app.Get("/about", h).Name = "about"
	app.Get("/page/{id}", h).Name = "page"

	fmt.Println(app.GetRoutes())

	app.Run(iris.Addr(":8080"))
}
