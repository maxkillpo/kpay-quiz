package model

import "github.com/globalsign/mgo/bson"

/*
## Merchant Fields
- Name
- Bank Account
- Username
- Password
*/

type Merchant struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Username     string        `bson:"username" json:"username"`
	Password     string        `bson:"password" json:"password"`
	BankAccounts []BankAccount `bson:"bank_account" json:"bank_account"`
	Products     []Product     `bson:"products" json:"products"`
}

type BankAccount struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	NumberAccount string        `bson:"number_account" json:"number_account"`
	Balance       float64       `bson:"balance" json:"balance"`
}

type Product struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Detail string        `bson:"detail" json:"detail"`
	Price  float64       `bson:"price" json:"price"`
}

type History struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
}
