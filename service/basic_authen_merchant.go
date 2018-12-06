package service

import (
	"kpay-quiz/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Authen struct {
	DB *store.DAO
}

func (a *Authen) BasicAuthenMerchant(c *gin.Context) {

	username, password, ok := c.Request.BasicAuth()
	if ok {
		if username == "" || password == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if merchants, err := a.DB.FindAll(); err == nil {
			for _, merchantsCollection := range *merchants {
				if merchantsCollection.Username == username && merchantsCollection.Password == password {
					return
				}
			}
			c.JSON(http.StatusUnauthorized, map[string]string{"result": "username or password not correct"})
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
