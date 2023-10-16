package model_test

import (
	"github.com/izakdvlpr/codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewBank(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"

	bank, err := model.NewBank(code, name)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(bank.ID))
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, err = model.NewBank("", "")

	require.NotNil(t, err)
}
