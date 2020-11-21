package main

import (
	"strconv"

	"github.com/kataras/iris/v12"
)

type Product struct {
	Id int
	Name string
	Price int
}

type Post struct {
	Id int
	Title string
}

type GetProductByIdResponse struct {
	IsFound bool
	ProductDetail Product 
}

type GetPostByIdResponse struct {
	IsFound bool
	PostDetail Post 
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

	app.Get("/products/{productID}", func(ctx iris.Context){
		productId := ctx.Params().Get("productID")
		productIdInt, _ := strconv.Atoi(productId)

		var response GetProductByIdResponse
		for _, item := range products {
			if productIdInt == item.Id {
				response.IsFound = true
				response.ProductDetail = item
			}
		}

		ctx.JSON(response)
	})
	app.Get("/posts", func(ctx iris.Context){
		ctx.JSON(posts)
	})

	app.Get("/posts/{postID}", func(ctx iris.Context){
		postId := ctx.Params().Get("postID")
		postIdInt, _ := strconv.Atoi(postId)

		var response GetPostByIdResponse
		for _, item := range posts {
			if postIdInt == item.Id {
				response.IsFound = true
				response.PostDetail = item
			}
		}

		ctx.JSON(response)
	})
	app.Listen(":8080")
}