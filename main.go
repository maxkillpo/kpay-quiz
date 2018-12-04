package main

import (
	"kpay-quiz/router"
	"kpay-quiz/service"
	"kpay-quiz/store"
	"os"
)

var (
	config = Config{}
	daos   = store.DAO{}
)

func init() {
	config.Read()
	daos.Server = config.Server
	daos.Database = config.Database
	daos.Connect()
}

func main() {

	s := &service.Server{
		Register: &service.RegisterServiceImprement{
			DB: &daos,
		},
		Merchant: &service.MerchantServiceImprement{
			DB: &daos,
		},
		MerchantProduct: &service.MerchantProductServiceImprement{
			DB: &daos,
		},
	}

	r := router.Setup(s)
	r.Run(":" + os.Getenv("PORT"))
}
