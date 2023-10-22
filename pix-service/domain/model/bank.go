package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Bank struct {
	Base `valid:"required"`

	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Code     string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Accounts []*Account `json:"accounts" gorm:"ForeignKey:BankID" valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}
