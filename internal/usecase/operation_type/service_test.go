package operation_type

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/internal/entity"
	"github.com/supwr/pismo-transactions/internal/usecase/operation_type/mock"
	"testing"
)

func Test_FindById(t *testing.T) {
	t.Run("find by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)

		operationType := &entity.OperationType{
			ID:   1,
			Name: "COMPRA A VISTA",
		}

		repo.EXPECT().FindById(operationType.ID).Return(operationType, nil).Times(1)

		service := NewService(repo)
		o, err := service.FindById(operationType.ID)

		assert.Equal(t, o, operationType)
		assert.Nil(t, err)
	})

	t.Run("error finding by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedError := errors.New("database error")

		repo.EXPECT().FindById(1).Return(nil, expectedError).Times(1)

		service := NewService(repo)
		o, err := service.FindById(1)

		assert.ErrorIs(t, err, expectedError)
		assert.Nil(t, o)
	})
}
