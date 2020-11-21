package main

import "github.com/kataras/iris/v12"

type Product struct {
	Id int
	Name string
	Price int
}

type Post struct {
	Id int
	Title string
}

var products = []Product{
	Product{
		Id: 1,
		Name: "Product 1",
		Price: 100,
	},
	Product{
		Id: 2,
		Name: "Product 2",
		Price: 200,
	},
}

var posts = []Post{
	Post{
		Id: 1,
		Title: "Post 1",
	},
	Post{
		Id: 2,
		Title: "Post 2",
	},
}

func main() {
	app := iris.Default()
	app.Get("/products", func(ctx iris.Context){
		ctx.JSON(products)
	})
	app.Get("/posts", func(ctx iris.Context){
		ctx.JSON(posts)
	})
	app.Listen(":8080")
}