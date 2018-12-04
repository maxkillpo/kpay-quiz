package service

import (
	"kpay-quiz/model"
	"kpay-quiz/store"
)

type MerchantService interface {
	Information(id string) (*model.Merchant, error)
}

type MerchantServiceImprement struct {
	DB *store.DAO
}

func (m *MerchantServiceImprement) Information(id string) (*model.Merchant, error) {
	merchant, err := m.DB.FindId(id)
	return merchant, err
}
