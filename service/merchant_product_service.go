package service

import (
	"kpay-quiz/model"
	"kpay-quiz/store"

	"github.com/globalsign/mgo/bson"
)

type MerchantProductService interface {
	Add(merchantId string, product *model.Product) (*model.Merchant, error)
	Remove(merchantId string, productId string) (*model.Merchant, error)
	Update(merchantId string, productId string, product *model.Product) (*model.Merchant, error)
	All(merchantId string) (*[]model.Product, error)
}

type MerchantProductServiceImprement struct {
	DB *store.DAO
}

func (mp *MerchantProductServiceImprement) Add(merchantId string, product *model.Product) (*model.Merchant, error) {
	merchant, err := mp.DB.FindMerchantById(merchantId)

	if err != nil {
		return merchant, err
	}
	product.ID = bson.NewObjectId()
	merchant.Products = append(merchant.Products, *product)

	return merchant, mp.DB.UpdateMerchantById(merchantId, merchant)
}

func (mp *MerchantProductServiceImprement) Update(merchantId string, productId string, product *model.Product) (*model.Merchant, error) {

	merchant, err := mp.DB.FindMerchantById(merchantId)

	if err != nil {
		return merchant, err
	}

	var products []model.Product

	for _, element := range merchant.Products {

		if element.ID == bson.ObjectIdHex(productId) {
			product.ID = element.ID
			products = append(products, *product)
		} else {
			products = append(products, element)
		}
	}

	merchant.Products = products

	return merchant, mp.DB.UpdateMerchantById(merchantId, merchant)
}

func (mp *MerchantProductServiceImprement) Remove(merchantId string, productId string) (*model.Merchant, error) {

	merchant, err := mp.DB.FindMerchantById(merchantId)

	if err != nil {
		return merchant, err
	}

	var products []model.Product

	for _, element := range merchant.Products {

		if element.ID != bson.ObjectIdHex(productId) {
			products = append(products, element)
		}
	}

	merchant.Products = products

	return merchant, mp.DB.UpdateMerchantById(merchantId, merchant)
}

func (mp *MerchantProductServiceImprement) All(merchantId string) (*[]model.Product, error) {
	merchant, err := mp.DB.FindMerchantById(merchantId)
	return &merchant.Products, err
}
