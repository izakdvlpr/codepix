package usecase

import (
	"errors"
	"github.com/izakdvlpr/codepix/domain/model"
	"github.com/izakdvlpr/codepix/infrastructure/repository"
)

type PixKeyUseCase struct {
	PixKeyRepository repository.PixKeyRepositoryInterface
}

func (u *PixKeyUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := u.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)

	if err != nil {
		return nil, err
	}

	u.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil
}

func (u *PixKeyUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := u.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
