package authen

import (
	"kpay-quiz/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Option struct {
	DB *store.DAO
}

func (a *Option) BasicAuthenMerchant(c *gin.Context) {

	if username, password, ok := c.Request.BasicAuth(); ok {
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
