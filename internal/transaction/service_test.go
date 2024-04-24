package transaction

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/internal/account"
	clockmock "github.com/supwr/pismo-transactions/pkg/clock/mock"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	t.Run("create cash buy transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		acc := &account.Account{
			ID:                   1,
			AvailableCreditLimit: decimal.NewFromInt(1000),
			Document:             "123456",
		}

		operationCashBuy := OperationTypeCashBuy

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(acc, nil).Times(1)

		transactionDate := time.Now()
		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: operationCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findAccountById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("create payment transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		acc := &account.Account{
			ID:                   1,
			Document:             "123456",
			AvailableCreditLimit: decimal.NewFromInt(1000),
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(acc, nil).Times(1)

		transactionDate := time.Now()
		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: OperationTypePayment,
			Amount:          decimal.NewFromFloat(float64(123.45)),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findAccountById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("find account by id error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		expectedError := errors.New("database error")
		transactionDate := time.Now()
		ctx := context.Background()

		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountRepo.EXPECT().FindById(ctx, 1).Return(nil, expectedError).Times(1)

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("account not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		transactionDate := time.Now()
		ctx := context.Background()

		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountRepo.EXPECT().FindById(ctx, 1).Return(nil, nil).Times(1)

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, ErrAccountNotFound)
	})

	t.Run("find operation type error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		expectedError := errors.New("database error")
		ctx := context.Background()

		acc := &account.Account{
			ID:                   1,
			Document:             "123456",
			AvailableCreditLimit: decimal.NewFromInt(1000),
		}

		accountRepo.EXPECT().FindById(ctx, 1).Return(acc, nil).Times(1)

		transactionDate := time.Now()
		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("operation type not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		acc := &account.Account{
			ID:                   1,
			Document:             "123456",
			AvailableCreditLimit: decimal.NewFromInt(1000),
		}

		accountRepo.EXPECT().FindById(ctx, 1).Return(acc, nil).Times(1)

		transactionDate := time.Now()
		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, ErrOperationTypeNotFound)
	})

	t.Run("create installment buy transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		acc := &account.Account{
			ID:                   1,
			Document:             "123456",
			AvailableCreditLimit: decimal.NewFromInt(1000),
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(acc, nil).Times(1)

		transactionDate := time.Now()
		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: OperationTypeInstallmentBuy,
			Amount:          decimal.NewFromFloat(float64(657.89)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findAccountById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("create withdraw transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := account.NewMockRepositoryInterface(ctrl)
		transactionRepo := NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		acc := &account.Account{
			ID:                   1,
			Document:             "123456",
			AvailableCreditLimit: decimal.NewFromInt(1000),
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(acc, nil).Times(1)

		transactionDate := time.Now()
		transaction := &Transaction{
			AccountID:       1,
			OperationTypeID: OperationTypeWithdraw,
			Amount:          decimal.NewFromFloat(float64(654.32)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findAccountById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := account.NewService(accountRepo)
		transactionService := NewService(transactionRepo, accountService, clockMock)

		err := transactionService.Create(ctx, &Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})
}
