package usecase

import (
	"errors"
	"github.com/izakdvlpr/codepix/domain/model"
	"github.com/izakdvlpr/codepix/infrastructure/repository"
	"log"
)

type TransactionUseCase struct {
	TransactionRepository repository.TransactionRepositoryInterface
	PixKeyRepository      repository.PixKeyRepositoryInterface
}

func (u *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string, id string) (*model.Transaction, error) {
	account, err := u.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := u.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	u.TransactionRepository.Save(transaction)

	if transaction.Base.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("unable to process this transaction")
}

func (u *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := u.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)

		return nil, err
	}

	transaction.Status = model.TransactionConfirmed

	err = u.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (u *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := u.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted

	err = u.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (u *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := u.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = u.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
