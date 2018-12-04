package service

import (
	"kpay-quiz/model"
	"kpay-quiz/store"
)

type MerchantService interface {
	Information(id string) (*model.Merchant, error)
	Update(id string, updateMerchantJSON *UpdateMerchantJSON) (*model.Merchant, error)
}

type MerchantServiceImprement struct {
	DB *store.DAO
}

func (m *MerchantServiceImprement) Information(id string) (*model.Merchant, error) {
	merchant, err := m.DB.FindMerchantById(id)
	return merchant, err
}

func (m *MerchantServiceImprement) Update(id string, updateMerchantJSON *UpdateMerchantJSON) (*model.Merchant, error) {

	merchant, err := m.DB.FindMerchantById(id)

	if err != nil {
		return merchant, err
	}

	if updateMerchantJSON.Name != "" {
		merchant.Name = updateMerchantJSON.Name
	}

	if updateMerchantJSON.BankAccountNumber != "" {
		merchant.BankAccount.NumberAccount = updateMerchantJSON.BankAccountNumber
	}

	m.DB.UpdateMerchantById(id, merchant)

	return merchant, err
}
