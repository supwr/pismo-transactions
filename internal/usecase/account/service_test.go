package account

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/internal/entity"
	"github.com/supwr/pismo-transactions/internal/usecase/account/mock"
	"testing"
	"time"
)

func Test_FindById(t *testing.T) {
	t.Run("find by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindById(account.ID).Return(account, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindById(account.ID)
		assert.Equal(t, account, a)
		assert.Nil(t, err)
	})

	t.Run("account not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)

		repo.EXPECT().FindById(1).Return(nil, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindById(1)
		assert.Nil(t, a)
		assert.Nil(t, err)
	})

	t.Run("error finding account by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedErr := errors.New("database error")

		repo.EXPECT().FindById(1).Return(nil, expectedErr).Times(1)

		service := NewService(repo)
		a, err := service.FindById(1)
		assert.Nil(t, a)
		assert.ErrorIs(t, err, expectedErr)
	})
}

func Test_FindByDocument(t *testing.T) {
	t.Run("find by document successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindByDocument(account.Document).Return(account, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindByDocument(account.Document)
		assert.Equal(t, account, a)
		assert.Nil(t, err)
	})

	t.Run("account not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		document := entity.Document("123456")

		repo.EXPECT().FindByDocument(document).Return(nil, nil).Times(1)

		service := NewService(repo)
		a, err := service.FindByDocument(document)
		assert.Nil(t, a)
		assert.Nil(t, err)
	})

	t.Run("error finding account by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		document := entity.Document("123456")
		expectedErr := errors.New("database error")

		repo.EXPECT().FindByDocument(document).Return(nil, expectedErr).Times(1)

		service := NewService(repo)
		a, err := service.FindByDocument(document)
		assert.Nil(t, a)
		assert.ErrorIs(t, err, expectedErr)
	})
}

func Test_Create(t *testing.T) {
	t.Run("create account successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		findByDocument := repo.EXPECT().FindByDocument(account.Document).Return(nil, nil).Times(1)
		repo.EXPECT().Create(account).Return(nil).Times(1).After(findByDocument)

		service := NewService(repo)
		err := service.Create(account)
		assert.Nil(t, err)
	})

	t.Run("error finding account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedErr := errors.New("database error")

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindByDocument(account.Document).Return(nil, expectedErr).Times(1)

		service := NewService(repo)
		err := service.Create(account)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("account already exists error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		repo.EXPECT().FindByDocument(account.Document).Return(account, nil).Times(1)

		service := NewService(repo)
		err := service.Create(account)
		assert.ErrorIs(t, err, ErrAccountAlreadyExists)
	})

	t.Run("error creating account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepositoryInterface(ctrl)
		expectedErr := errors.New("database error")

		account := &entity.Account{
			ID:        1,
			Document:  "123456",
			CreatedAt: time.Now(),
		}

		findByDocument := repo.EXPECT().FindByDocument(account.Document).Return(nil, nil).Times(1)
		repo.EXPECT().Create(account).Return(expectedErr).Times(1).After(findByDocument)

		service := NewService(repo)
		err := service.Create(account)
		assert.ErrorIs(t, err, expectedErr)
	})
}
