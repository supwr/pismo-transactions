package transaction

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/entity"
	clockmock "github.com/supwr/pismo-transactions/pkg/clock/mock"
	accountservice "github.com/supwr/pismo-transactions/usecase/account"
	accountrepo "github.com/supwr/pismo-transactions/usecase/account/mock"
	operationtypeservice "github.com/supwr/pismo-transactions/usecase/operation_type"
	operationtyperepo "github.com/supwr/pismo-transactions/usecase/operation_type/mock"
	transactionrepo "github.com/supwr/pismo-transactions/usecase/transaction/mock"
	"testing"
	"time"
)

func Test_Create(t *testing.T) {
	t.Run("create cash buy transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationType := &entity.OperationType{
			ID:   entity.OperationTypeCashBuy,
			Name: "COMPRA A VISTA",
		}

		findAccountById := accountRepo.EXPECT().FindById(1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(entity.OperationTypeCashBuy).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("create payment transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationType := &entity.OperationType{
			ID:   entity.OperationTypePayment,
			Name: "PAGAMENTO",
		}

		findAccountById := accountRepo.EXPECT().FindById(1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(entity.OperationTypePayment).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypePayment,
			Amount:          decimal.NewFromFloat(float64(123.45)),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("find account by id error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		expectedError := errors.New("database error")
		transactionDate := time.Now()

		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountRepo.EXPECT().FindById(1).Return(nil, expectedError).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("account not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		transactionDate := time.Now()

		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountRepo.EXPECT().FindById(1).Return(nil, nil).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, ErrAccountNotFound)
	})

	t.Run("find operation type error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		expectedError := errors.New("database error")

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		findAccountById := accountRepo.EXPECT().FindById(1).Return(account, nil).Times(1)
		operationTypeRepo.EXPECT().FindById(entity.OperationTypeCashBuy).Return(nil, expectedError).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("operation type not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		findAccountById := accountRepo.EXPECT().FindById(1).Return(account, nil).Times(1)
		operationTypeRepo.EXPECT().FindById(entity.OperationTypeCashBuy).Return(nil, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, ErrOperationTypeNotFound)
	})

	t.Run("create installment buy transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationType := &entity.OperationType{
			ID:   entity.OperationTypeInstallmentBuy,
			Name: "COMPRA PARCELADA",
		}

		findAccountById := accountRepo.EXPECT().FindById(1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(entity.OperationTypeInstallmentBuy).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeInstallmentBuy,
			Amount:          decimal.NewFromFloat(float64(657.89)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("create withdraw transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationType := &entity.OperationType{
			ID:   entity.OperationTypeWithdraw,
			Name: "SAQUE",
		}

		findAccountById := accountRepo.EXPECT().FindById(1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(entity.OperationTypeWithdraw).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeWithdraw,
			Amount:          decimal.NewFromFloat(float64(654.32)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})
}
