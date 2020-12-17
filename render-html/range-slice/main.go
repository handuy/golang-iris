package main

import "github.com/kataras/iris/v12"

type Product struct {
	Img  string
	Name string
	Price int
	Sold int
}

type News struct {
	Img string
	Title string
	Description string
	Date string
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

	app.Get("/products", func(ctx iris.Context){
		products := []Product{
			Product{
				Img: "/resources/image/product/adidas/Adidas-Yeezy-Boost-350-V2-Core-Black-Red-Product.jpg",
				Name: "Adidas Yeezy Boost 350 V2 Core Black Red",
				Price: 12015122,
				Sold: 2000,
			},
			Product{
				Img: "/resources/image/product/nike/Nike-Air-Max-90-Off-White-Product.jpg",
				Name: "Air Max 90 OFF-WHITE",
				Price: 4810000,
				Sold: 1312,
			},
			Product{
				Img: "/resources/image/product/nike/Nike-Cortez-Stranger-Things-Independence-Day-Pack.jpeg",
				Name: "Nike Classic Cortez Stranger Things Independence Day Pack",
				Price: 20790000,
				Sold: 2246,
			},
			Product{
				Img: "/resources/image/product/nike/Nike-Air-Max-1-97-Sean-Wotherspoon-NA-Product.jpg",
				Name: "Air Max 1/97 Sean Wotherspoon",
				Price: 3740000,
				Sold: 564,
			},
			Product{
				Img: "/resources/image/product/nike/Nike-Air-Force-1-Low-Off-White-Black-White-Product.jpg",
				Name: "Air Force 1 Low Off-White Black White",
				Price: 6620000,
				Sold: 1456,
			},
			Product{
				Img: "/resources/image/product/nike/Nike-Air-Presto-Off-White-White-2018-Product.jpg",
				Name: "Air Presto Off-White White (2018)",
				Price: 10000000,
				Sold: 14560,
			},
			Product{
				Img: "/resources/image/product/nike/Nike-Blazer-High-sacai-Snow-Beach-Product.jpg",
				Name: "Nike Blazer High sacai Snow Beach",
				Price: 1030000,
				Sold: 14520,
			},
			Product{
				Img: "/resources/image/product/nike/Nike-Adapt-BB-Mag-Product.jpg",
				Name: "Nike Adapt BB Mag (Sạc bản US)",
				Price: 10040000,
				Sold: 14530,
			},
			Product{
				Img: "/resources/image/product/converse/Converse-Chuck-Taylor-All-Star-70s-Hi-Brain-Dead-Product.jpg",
				Name: "Converse Chuck Taylor All-Star 70s Hi Brain Dead",
				Price: 10000000,
				Sold: 14560,
			},
			Product{
				Img: "/resources/image/product/adidas/Adidas-Human-Race-NMD-Pharrell-x-Chanel-Black.jpeg",
				Name: "Adidas Human Race NMD Pharrell x Chanel",
				Price: 10000000,
				Sold: 24560,
			},
		}
		
		ctx.ViewData("products", products)
		ctx.View("products/product-after.html")
	})

	app.Get("/news", func(ctx iris.Context){
		newsList := []News{
			News{
				Img: "/resources/image/news/news-1-thumnails.jpg",
				Title: "Pack “Logo Sketch” với branding mộc mạc trên Air Max 97 và Air Max Plus",
				Description: "Trong pack này, chúng ta sẽ đến với mẫu branding kiểu phác thảo, đầy mộc mạc và “blend” hoàn hảo cùng tổng thể phối màu.",
				Date: "20/1/2020",
			},
			News{
				Img: "/resources/image/news/news-2-thumnails.jpg",
				Title: "Nike Classic Cortez – Phối Màu Của Sự Tinh Tế",
				Description: "Phối màu tiếp theo gia nhập bộ sưu tập Nike Classic Cortez với sự kết hợp của 2 tông màu: Hồng Hoa Đăng và màu Cam",
				Date: "20/1/2020",
			},
			News{
				Img: "/resources/image/news/news-3-thumnails.jpg",
				Title: "Air Jordan 1 “Satin Black Toe” xuất hiện hình ảnh chi tiết",
				Description: "Air Jordan 1 “Black Toe” vẫn luôn là thiết kế Air Jordan 1 đáng mua nhất mọi thời đại. Vậy sẽ thế nào nếu chúng khoác lên mình chất liệu satin?",
				Date: "20/1/2020",
			},
			News{
				Img: "/resources/image/news/news-4-thumnails.jpg",
				Title: "HOT! Nike và Levi’s tái hợp với hàng loạt siêu phẩm mới",
				Description: "Đến hẹn lại lên, Nike và Levi’s lại tái hợp với nhau khiến các sneakerheads và tín đồ thời trang mê mệt.",
				Date: "20/1/2020",
			},
			News{
				Img: "/resources/image/news/news-5-thumnails.jpg",
				Title: "Xuất hiện phối màu thứ 3 của Nike Air Force 1 Type N354",
				Description: "Nike N354 là tên dự án với những thiết kế lấy cảm hứng từ kho Archive khổng lồ của The Swoosh. Nổi bật trong số đó là phiên bản Nike Air Force 1 Type.",
				Date: "20/1/2020",
			},
		}
		
		ctx.ViewData("newsList", newsList)
		ctx.View("news/news.html")
	})

	app.Listen(":8080")
}