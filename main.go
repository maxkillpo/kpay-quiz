package main

import (
	"kpay-quiz/store"
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

	o := store.Product{
		Id:     "1",
		Name:   "Max",
		Detail: "AAAA",
		Price:  12,
	}

	daos.Insert(o)
}
