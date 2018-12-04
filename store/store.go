package store

import (
	"kpay-quiz/model"
	"log"

	"github.com/globalsign/mgo"
)

type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "Merchant"

func (d *DAO) Connect() {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
}

func (d *DAO) Insert(product model.Product) error {
	return db.C(COLLECTION).Insert(&product)
}
