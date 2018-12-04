package store

import (
	"kpay-quiz/model"
	"log"

	"github.com/globalsign/mgo/bson"

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

func (d *DAO) InsertMerchant(any interface{}) error {
	return db.C(COLLECTION).Insert(&any)
}

func (d *DAO) FindMerchantById(id string) (*model.Merchant, error) {
	var merchart model.Merchant
	return &merchart, db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&merchart)
}

func (d *DAO) UpdateMerchantById(id string, merchant *model.Merchant) error {
	return db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), merchant)
}
