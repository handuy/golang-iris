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
		products := []Product{
			Product{
				Name: "Adidas Yeezy Boost 350 V2 Lundmark 123",
				Price: 9000000,
				Sold: 2000,
			},
			Product{
				Name: "Jordan 4 Retro Cool Grey (2019)",
				Price: 4810000,
				Sold: 1312,
			},
			Product{
				Name: "Jordan 1 Retro Low OG SP Travis Scott",
				Price: 20790000,
				Sold: 2246,
			},
			Product{
				Name: "Air Jordan 1 Mid Yellow Toe Black",
				Price: 3740000,
				Sold: 564,
			},
			Product{
				Name: "Nike SB Dunk Low Parra",
				Price: 6620000,
				Sold: 1456,
			},
			Product{
				Name: "Adidas Limited Edition",
				Price: 10000000,
				Sold: 14560,
			},
		}

		ctx.ViewData("list", products)
		ctx.View("home-page/index.html")
	})

	app.Listen(":8080")
}