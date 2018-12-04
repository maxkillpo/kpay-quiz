package service

import (
	"errors"
	model "kpay-quiz/model"
	"kpay-quiz/store"
)

type RegisterService interface {
	Merchant(merchant *model.Merchant) error
}

type RegisterServiceImprement struct {
	DB *store.DAO
}

func (r *RegisterServiceImprement) Merchant(merchant *model.Merchant) error {
	if merchant.Id != "" && merchant.BankAccount != nil && merchant.Username != "" && merchant.Password != "" {
		return errors.New("connot have: Id & BankAccount & Username & Password")
	}

	return nil
}
