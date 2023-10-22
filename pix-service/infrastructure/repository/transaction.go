package repository

import (
	"fmt"
	"github.com/izakdvlpr/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryInterface interface {
	Register(transaction *model.Transaction) error
	Save(transaction *model.Transaction) error
	Find(id string) (*model.Transaction, error)
}

type TransactionRepositoryDatabase struct {
	Database *gorm.DB
}

func (r *TransactionRepositoryDatabase) Register(transaction *model.Transaction) error {
	err := r.Database.Create(transaction).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryDatabase) Save(transaction *model.Transaction) error {
	err := r.Database.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDatabase) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Database.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
