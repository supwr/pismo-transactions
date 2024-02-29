package transaction

import (
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/entity"
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

		transactionRepo.EXPECT().Create(transaction).Return(nil).After(findOperationTypeById).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService)

		err := transactionService.Create(&entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})
}
