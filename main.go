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

	// o := store.Product{
	// 	Id:     "1",
	// 	Name:   "Max",
	// 	Detail: "AAAA",
	// 	Price:  12,
	// }

	// daos.Insert(o)

	s := &service.Server{
		Register: &service.RegisterServiceImprement{
			DB: &daos,
		},
	}

	r := router.Setup(s)
	r.Run(":" + os.Getenv("PORT"))
}
