package main

import (
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

type Product struct {
	Id    int
	Name  string
	Price int
}

type Post struct {
	Id    int
	Title string
}

type GetProductByIdResponse struct {
	IsFound       bool
	ProductDetail Product
}

type GetPostByIdResponse struct {
	IsFound    bool
	PostDetail Post
}

var products = []Product{
	Product{
		Id:    1,
		Name:  "Real Madrid",
		Price: 100,
	},
	Product{
		Id:    2,
		Name:  "Atletico Madrid",
		Price: 200,
	},
	Product{
		Id:    3,
		Name:  "Man United",
		Price: 100,
	},
	Product{
		Id:    4,
		Name:  "Newcastle United",
		Price: 200,
	},
}

var posts = []Post{
	Post{
		Id:    1,
		Title: "Golang",
	},
	Post{
		Id:    2,
		Title: "Java",
	},
	Post{
		Id:    3,
		Title: "Javascript",
	},
	Post{
		Id:    4,
		Title: "Python",
	},
}

func main() {
	app := iris.Default()

	// /products?keyword=hi
	app.Get("/products", func(ctx iris.Context) {
		searchReq := ctx.URLParam("keyword")

		var searchRsp []Product

		if searchReq != "" {
			for _, item := range products {
				if strings.Contains(item.Name, searchReq) {
					searchRsp = append(searchRsp, item)
				}
			}
		} else {
			searchRsp = products
		}

		ctx.JSON(searchRsp)
	})

	app.Get("/products/{productID}", func(ctx iris.Context) {
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

	app.Post("/products/create", func(ctx iris.Context) {
		productName := ctx.PostValue("name")
		productPrice := ctx.PostValue("price")
		productPriceInt, _ := strconv.Atoi(productPrice)
		productId := len(products) + 1

		newProduct := Product{
			Id:    productId,
			Name:  productName,
			Price: productPriceInt,
		}

		products = append(products, newProduct)
	})

	app.Get("/posts", func(ctx iris.Context) {
		searchReq := ctx.URLParam("keyword")

		var searchRsp []Post

		if searchReq != "" {
			for _, item := range posts {
				if strings.Contains(item.Title, searchReq) {
					searchRsp = append(searchRsp, item)
				}
			}
		} else {
			searchRsp = posts
		}

		ctx.JSON(searchRsp)
	})

	app.Get("/posts/{postID}", func(ctx iris.Context) {
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

	app.Post("/posts/create", func(ctx iris.Context) {
		postTitle := ctx.PostValue("title")
		postId := len(posts) + 1

		newPost := Post{
			Id:    postId,
			Title:  postTitle,
		}

		posts = append(posts, newPost)
	})

	app.Listen(":8080")
}
