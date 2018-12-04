package service

import (
	"fmt"
	"kpay-quiz/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Register RegisterService
}

func (s *Server) RegisterMerchant(c *gin.Context) {

	var merchant model.Merchant
	err := c.ShouldBindJSON(&merchant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("json: wrong params: %s", err),
		})
		return
	}

	if err := s.Register.Merchant(&merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, merchant)
}
