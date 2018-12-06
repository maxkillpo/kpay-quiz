package router

import (
	"kpay-quiz/service"
	"kpay-quiz/store"

	"github.com/gin-gonic/gin"
)

/*
| Name                 | Method | Endpoint                          |
|----------------------|--------|-----------------------------------|
| Register Merchant    | POST   | /register                         |
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
	setupRouterRegister(s, router)
	setupRouterMerchant(s, router)
	setupRouterBuy(s, router)
	return router
}

func setupRouterRegister(s *service.Server, g *gin.Engine) {
	g.POST("/register", s.RegisterMerchant)
}

func setupRouterMerchant(s *service.Server, g *gin.Engine) {
	authen := service.Authen{
		DB: &store.DAOS,
	}
	merchant := g.Group("/merchant")
	merchant.Use(authen.BasicAuthenMerchant)
	merchant.GET("/:id", s.MerchantInformation)
	merchant.POST("/:id", s.UpdateMerchant)
	merchant.GET("/:id/products", s.ListAllProducts)
	merchant.POST("/:id/product", s.AddProduct)
	merchant.POST("/:id/product/:product_id", s.UpdateProduct)
	merchant.DELETE("/:id/product/:product_id", s.RemoveProduct)
	merchant.POST("/:id/report", s.SellReports)
}

func setupRouterBuy(s *service.Server, g *gin.Engine) {
	buy := g.Group("/buy")
	buy.POST("/product", s.BuyProduct)
}
