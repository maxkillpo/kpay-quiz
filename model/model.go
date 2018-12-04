package model

import "gopkg.in/mgo.v2/bson"

/*
## Merchant Fields
- Name
- Bank Account
- Username
- Password
*/

type Merchant struct {
	Id          bson.ObjectId `bson: "_id" json:"id"`
	Name        string        `bson: "name" json:"name"`
	BankAccount []BankAccount `bson: "bank_account" json:"bank_account"`
	Username    string        `bson: "username" json:"username"`
	Password    string        `bson: "password" json:"password"`
}

type BankAccount struct {
	Id            bson.ObjectId  `bson: "_id" json:"id"`
	NumberAccount string  `bson: "number_account" json:"number_account"`
	Balance       float64 `bson: "balance" json:"balance"`
}

type Product struct {
	Id     bson.ObjectId  `bson: "_id" json:"id"`
	Name   string  `bson: "name" json:"name"`
	Detail string  `bson: "detail" json:"detail"`
	Price  float64 `bson: "price" json:"price"`
}
