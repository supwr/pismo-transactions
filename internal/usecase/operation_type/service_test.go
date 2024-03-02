package operation_type

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/internal/entity"
	"github.com/supwr/pismo-transactions/internal/usecase/operation_type/mock"
	"testing"
)

func TestService_FindById(t *testing.T) {
	t.Run("find by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		ctx := context.Background()

		operationType := &entity.OperationType{
			ID:   1,
			Name: "COMPRA A VISTA",
		}

		repo.EXPECT().FindById(ctx, operationType.ID).Return(operationType, nil).Times(1)

		service := NewService(repo)
		o, err := service.FindById(ctx, operationType.ID)

		assert.Equal(t, o, operationType)
		assert.Nil(t, err)
	})

	t.Run("error finding by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedError := errors.New("database error")
		ctx := context.Background()

		repo.EXPECT().FindById(ctx, 1).Return(nil, expectedError).Times(1)

		service := NewService(repo)
		o, err := service.FindById(ctx, 1)

		assert.ErrorIs(t, err, expectedError)
		assert.Nil(t, o)
	})
}
