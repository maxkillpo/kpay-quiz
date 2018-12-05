package service

import (
	"encoding/json"
	"fmt"
	"kpay-quiz/model"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var mutex sync.Mutex

type Server struct {
	Register        RegisterService
	Merchant        MerchantService
	MerchantProduct MerchantProductService
	Buy             BuyService
}

type RegisterJSON struct {
	Name              string `bson:"name" json:"name"`
	BankAccountNumber string `bson:"bank_account_number" json:"bank_account_number"`
}

func (s *Server) RegisterMerchant(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()
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

	printLog(merchant)

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
	mutex.Lock()
	defer mutex.Unlock()

	id := c.Param("id")
	merchant, err := s.Merchant.Information(id)

	printLog(merchant)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	data := map[string]interface{}{
		"data": merchant,
	}

	c.JSON(http.StatusCreated, data)
}

type UpdateMerchantJSON struct {
	Name              string `bson:"name" json:"name"`
	BankAccountNumber string `bson:"bank_account_number" json:"bank_account_number"`
}

func (s *Server) UpdateMerchant(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	id := c.Param("id")
	var updateMerchantJSON UpdateMerchantJSON
	err := c.ShouldBindJSON(&updateMerchantJSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchant, err := s.Merchant.Update(id, &updateMerchantJSON)

	printLog(merchant)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (s *Server) AddProduct(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	id := c.Param("id")
	var product model.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchant, err := s.MerchantProduct.Add(id, &product)

	printLog(merchant)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (s *Server) ListAllProducts(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	id := c.Param("id")
	products, err := s.MerchantProduct.All(id)

	printLog(products)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	c.JSON(http.StatusCreated, products)

}

func (s *Server) UpdateProduct(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	var product model.Product
	merchectId := c.Param("id")
	productId := c.Param("product_id")

	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchent, err := s.MerchantProduct.Update(merchectId, productId, &product)

	printLog(merchent)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (s *Server) RemoveProduct(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	merchectId := c.Param("id")
	productId := c.Param("product_id")

	merchent, err := s.MerchantProduct.Remove(merchectId, productId)

	printLog(merchent)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (s *Server) SellReports(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

}

type BuyProductRequest struct {
	MerchectId       string `bson:"merchect_id" json:"merchect_id"`
	ProductId        string `bson:"product_id" json:"product_id"`
	BuyAmountProduct int    `bson:"buy_amount_product" json:"buy_amount_product"`
}

func (s *Server) BuyProduct(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	var buyProductRequest BuyProductRequest
	err := c.ShouldBindJSON(&buyProductRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	merchent, err := s.Buy.Product(&buyProductRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	printLog(merchent)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": merchent.BankAccount,
	})
}

func printLog(n interface{}) {
	b, _ := json.MarshalIndent(n, "", "\t")
	os.Stdout.Write(b)
}
