package store

type Product struct {
	Id     string  `bson: "_id" json:"id"`
	Name   string  `bson: "name" json:"name"`
	Detail string  `bson: "detail" json:"detail"`
	Price  float64 `bson: "price" json:"price"`
}
