package service

import (
	"errors"
	"kpay-quiz/model"
	"kpay-quiz/store"
	"math/rand"

	"github.com/globalsign/mgo/bson"
)

type RegisterService interface {
	Merchant(registerJSON *RegisterJSON) (*model.Merchant, error)
}

type RegisterServiceImprement struct {
	DB *store.DAO
}

func (r *RegisterServiceImprement) Merchant(registerJSON *RegisterJSON) (*model.Merchant, error) {

	if registerJSON.Name == "" {
		return nil, errors.New("found name")
	}

	if registerJSON.BankAccountNumber == "" {
		return nil, errors.New("found bank_account_number")
	}

	textForRandomUsername := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	textForRandomPassword := "0123456789@#!_-+&$M?"

	username := make([]byte, 10)
	for i := range username {
		username[i] = textForRandomUsername[rand.Intn(len(textForRandomUsername))]
	}

	password := make([]byte, 10)
	for i := range password {
		password[i] = textForRandomPassword[rand.Intn(len(textForRandomPassword))]
	}

	createBankAccount := []model.BankAccount{}

	createBankAccount = append(createBankAccount, model.BankAccount{
		ID:            bson.NewObjectId(),
		NumberAccount: registerJSON.BankAccountNumber,
		Balance:       0.0,
	})

	merchant := model.Merchant{
		ID:           bson.NewObjectId(),
		Name:         registerJSON.Name,
		Username:     string(username),
		Password:     string(password),
		BankAccounts: createBankAccount,
		Products:     nil,
	}

	return &merchant, r.DB.InsertMerchant(merchant)
}
