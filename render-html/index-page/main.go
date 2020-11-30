package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context){
		ctx.View("home-page/index.html")
	})

	app.Listen(":8080")
}