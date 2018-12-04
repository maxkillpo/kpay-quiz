package router

import (
	"kpay-quiz/service"

	"github.com/gin-gonic/gin"
)

/*
| Name                 | Method | Endpoint                          |
|----------------------|--------|-----------------------------------|
| Register Merchant    | POST   | /merchant/register                |
| Merchant Information | GET    | /merchant/:id                     |
| Update Merchant      | POST   | /merchant/:id                     |
| List All Products    | GET    | /merchant/:id/products            |
| Add Product          | POST   | /merchant/:id/product             |
| Update Product       | POST   | /merchant/:id/product/:product_id |
| Remove Product       | DELETE | /merchant/:id/product/:product_id |
| Sell Reports         | POST   | /merchant/:id/report              |
| Buy Product          | POST   | /buy/product                      |
*/

func Setup(s *service.Server) *gin.Engine {
	router := gin.Default()
	setupRouterMerchant(s, router)
	setupRouterBuy(s, router)
	return router
}

func setupRouterMerchant(s *service.Server, g *gin.Engine) {
	merchant := g.Group("/merchant")
	merchant.POST("/register", s.RegisterMerchant)
	merchant.GET("/information/:id", s.MerchantInformation)
	merchant.POST("/update/:id", s.UpdateMerchant)
	product := merchant.Group("/product")
	product.GET("/:id/products", s.ListAllProducts)
	product.POST("/:id/product", s.AddProduct)
	product.POST("/:id/product/:product_id", s.UpdateProduct)
	product.DELETE("/:id/product/:product_id", s.RemoveProduct)
	report := merchant.Group("/report")
	report.POST("/:id/report", s.SellReports)
}

func setupRouterBuy(s *service.Server, g *gin.Engine) {
	buy := g.Group("/buy")
	buy.POST("/product", s.BuyProduct)
}
