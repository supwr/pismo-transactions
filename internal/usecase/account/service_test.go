package account

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/internal/entity"
	"github.com/supwr/pismo-transactions/internal/usecase/account/mock"
	"testing"
	"time"
)

func TestService_FindById(t *testing.T) {
	t.Run("find by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindById(ctx, account.ID).Return(account, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindById(ctx, account.ID)
		assert.Equal(t, account, a)
		assert.Nil(t, err)
	})

	t.Run("account not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		ctx := context.Background()

		repo.EXPECT().FindById(ctx, 1).Return(nil, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindById(ctx, 1)
		assert.Nil(t, a)
		assert.Nil(t, err)
	})

	t.Run("error finding account by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedErr := errors.New("database error")
		ctx := context.Background()

		repo.EXPECT().FindById(ctx, 1).Return(nil, expectedErr).Times(1)

		service := NewService(repo)
		a, err := service.FindById(ctx, 1)
		assert.Nil(t, a)
		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestService_FindByDocument(t *testing.T) {
	t.Run("find by document successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindByDocument(ctx, account.Document).Return(account, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindByDocument(ctx, account.Document)
		assert.Equal(t, account, a)
		assert.Nil(t, err)
	})

	t.Run("account not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		document := entity.Document("123456")
		ctx := context.Background()

		repo.EXPECT().FindByDocument(ctx, document).Return(nil, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindByDocument(ctx, document)
		assert.Nil(t, a)
		assert.Nil(t, err)
	})

	t.Run("error finding account by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		document := entity.Document("123456")
		expectedErr := errors.New("database error")
		ctx := context.Background()

		repo.EXPECT().FindByDocument(ctx, document).Return(nil, expectedErr).Times(1)

		service := NewService(repo)
		a, err := service.FindByDocument(ctx, document)
		assert.Nil(t, a)
		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestService_Create(t *testing.T) {
	t.Run("create account successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		findByDocument := repo.EXPECT().FindByDocument(ctx, account.Document).Return(nil, nil).Times(1)
		repo.EXPECT().Create(ctx, account).Return(nil).Times(1).After(findByDocument)

		service := NewService(repo)
		err := service.Create(ctx, account)
		assert.Nil(t, err)
	})

	t.Run("error finding account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedErr := errors.New("database error")
		ctx := context.Background()

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindByDocument(ctx, account.Document).Return(nil, expectedErr).Times(1)

		service := NewService(repo)
		err := service.Create(ctx, account)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("account already exists error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindByDocument(ctx, account.Document).Return(account, nil).Times(1)

		service := NewService(repo)
		err := service.Create(ctx, account)
		assert.ErrorIs(t, err, ErrAccountAlreadyExists)
	})

	t.Run("error creating account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedErr := errors.New("database error")
		ctx := context.Background()

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		findByDocument := repo.EXPECT().FindByDocument(ctx, account.Document).Return(nil, nil).Times(1)
		repo.EXPECT().Create(ctx, account).Return(expectedErr).Times(1).After(findByDocument)

		service := NewService(repo)
		err := service.Create(ctx, account)
		assert.ErrorIs(t, err, expectedErr)
	})
}
