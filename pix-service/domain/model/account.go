package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base `valid:"required"`

	Bank      *Bank     `json:"bank" valid:"-"`
	BankID    string    `json:"bank_id" gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	OwnerName string    `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	PixKeys   []*PixKey `json:"pixes" gorm:"ForeignKey:AccountID" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		Bank:      bank,
		Number:    number,
		OwnerName: ownerName,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &account, nil
}
