package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context){
		ctx.HTML("Hello World")
	})
	app.Listen(":8080")
}