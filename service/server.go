package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Register RegisterService
	Merchant MerchantService
}

type RegisterJSON struct {
	Name              string `bson:"name" json:"name"`
	BankAccountNumber string `bson:"bank_account_number" json:"bank_account_number"`
}

/*
{
	"name" : "qefqefqef",
	"bank_account_number" : "111-222-444-5-5-552"
}
*/

func (s *Server) RegisterMerchant(c *gin.Context) {
	var registerJSON RegisterJSON
	err := c.ShouldBindJSON(&registerJSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchant, err := s.Register.Merchant(&registerJSON)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	c.JSON(http.StatusCreated, merchant)
}

func (s *Server) MerchantInformation(c *gin.Context) {
	id := c.Param("id")
	merchant, err := s.Merchant.Information(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	c.JSON(http.StatusCreated, merchant)
}

func (s *Server) UpdateMerchant(c *gin.Context) {

}

func (s *Server) ListAllProducts(c *gin.Context) {

}

func (s *Server) AddProduct(c *gin.Context) {

}

func (s *Server) UpdateProduct(c *gin.Context) {

}

func (s *Server) RemoveProduct(c *gin.Context) {

}

func (s *Server) SellReports(c *gin.Context) {

}

func (s *Server) BuyProduct(c *gin.Context) {

}
