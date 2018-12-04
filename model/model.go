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
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Username    string        `bson:"username" json:"username"`
	Password    string        `bson:"password" json:"password"`
	BankAccount BankAccount   `bson:"bank_account" json:"bank_account"`
	Products    []Product     `bson:"products" json:"products"`
	History     []History     `bson:"history" json:"history"`
}

type BankAccount struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	NumberAccount string        `bson:"number_account" json:"number_account"`
	Balance       float64       `bson:"balance" json:"balance"`
}

type Product struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Price  float64       `bson:"price" json:"price"`
	Amount int           `bson:"amount" json:"amount"`
}

type History struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Amount      int           `bson:"amount" json:"amount"`
	ProductName string        `bson:"product_name" json:"product_name"`
	Date        string        `bson:"date" json:"date"`
}
