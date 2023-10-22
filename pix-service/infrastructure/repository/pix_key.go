package repository

import (
	"fmt"
	"github.com/izakdvlpr/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *model.PixKey) (*model.PixKey, error)
	FindKeyByKind(key string, kind string) (*model.PixKey, error)
	AddBank(bank *model.Bank) error
	AddAccount(account *model.Account) error
	FindAccount(id string) (*model.Account, error)
}

type PixKeyRepositoryDatabase struct {
	Database *gorm.DB
}

func (r PixKeyRepositoryDatabase) AddBank(bank *model.Bank) error {
	err := r.Database.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDatabase) AddAccount(account *model.Account) error {
	err := r.Database.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDatabase) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Database.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (r PixKeyRepositoryDatabase) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Database.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

func (r PixKeyRepositoryDatabase) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Database.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}

	return &account, nil
}

func (r PixKeyRepositoryDatabase) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Database.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}

	return &bank, nil
}
