package main

import (
	"kpay-quiz/router"
	"kpay-quiz/service"
	"kpay-quiz/store"
	"os"
)

func main() {

	s := &service.Server{
		Register: &service.RegisterServiceImprement{
			DB: &store.DAOS,
		},
		Merchant: &service.MerchantServiceImprement{
			DB: &store.DAOS,
		},
		MerchantProduct: &service.MerchantProductServiceImprement{
			DB: &store.DAOS,
		},
		Buy: &service.BuyServiceImprement{
			DB: &store.DAOS,
		},
		Report: &service.SellReportsServiceImprement{
			DB: &store.DAOS,
		},
	}

	r := router.Setup(s)
	r.Run(":" + os.Getenv("PORT"))
}
