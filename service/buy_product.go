package service

import (
	"errors"
	"kpay-quiz/model"
	"kpay-quiz/store"
	"time"

	"github.com/globalsign/mgo/bson"
)

type BuyService interface {
	Product(buy *BuyProductRequest) (*model.Merchant, error)
}

type BuyServiceImprement struct {
	DB *store.DAO
}

func (b *BuyServiceImprement) Product(buy *BuyProductRequest) (*model.Merchant, error) {

	merchantId := buy.MerchectId
	productId := buy.ProductId
	amountProduct := buy.BuyAmountProduct

	merchant, err := b.DB.FindMerchantById(merchantId)

	if err != nil {
		return merchant, err
	}

	for index, element := range merchant.Products {

		if element.ID == bson.ObjectIdHex(productId) {
			merchant.Products[index].Amount = element.Amount - amountProduct

			if merchant.Products[index].Amount < 0 {
				return merchant, errors.New("Product amount less than.")
			}

			totlePrice := (merchant.Products[index].Price * float64(amountProduct))
			merchant.BankAccount.Balance = merchant.BankAccount.Balance + totlePrice
			merchant.History = append(merchant.History, model.History{
				ID:          bson.NewObjectId(),
				Amount:      amountProduct,
				ProductName: element.Name,
				TotalPrice:  totlePrice,
				Date:        time.Now().String(),
			})
			break
		}
	}

	return merchant, b.DB.UpdateMerchantById(merchantId, merchant)
}
