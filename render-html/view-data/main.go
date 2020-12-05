package main

import "github.com/kataras/iris/v12"

type Product struct {
	Name string
	Price int
	Sold int
}

func main() {
	app := iris.Default()
	tmpl := iris.HTML("./view", ".html")
	app.RegisterView(tmpl)

	app.HandleDir("/resources", "./resources")

	app.Get("/", func(ctx iris.Context){
		product := Product{
			Name: "Adidas Yeezy Boost 350 V2 Lundmark 123",
			Price: 9000000,
			Sold: 2000,
		}

		ctx.ViewData("nike", product)
		ctx.View("home-page/index.html")
	})

	app.Listen(":8080")
}