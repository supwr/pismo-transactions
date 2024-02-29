package transaction

import (
	"github.com/supwr/pismo-transactions/entity"
	"github.com/supwr/pismo-transactions/usecase/account"
	"github.com/supwr/pismo-transactions/usecase/operation_type"
	"slices"
	"time"
)

type Service struct {
	repository           RepositoryInterface
	operationTypeService *operation_type.Service
	accountService       *account.Service
}

func NewService(r RepositoryInterface, o *operation_type.Service, a *account.Service) *Service {
	return &Service{repository: r, operationTypeService: o, accountService: a}
}

func (s *Service) Create(t *entity.Transaction) error {
	var negAmountTransactions = []int{entity.OperationTypeCashBuy, entity.OperationTypeInstallmentBuy, entity.OperationTypeWithdraw}

	acc, err := s.accountService.FindById(t.AccountID)
	if err != nil {
		return err
	}

	if acc == nil {
		return ErrAccountNotFound
	}

	operationType, err := s.operationTypeService.FindById(t.OperationTypeID)
	if err != nil {
		return err
	}

	if operationType == nil {
		return ErrOperationTypeNotFound
	}

	if slices.Contains(negAmountTransactions, t.OperationTypeID) {
		t.Amount = t.Amount.Abs().Neg()
	} else {
		t.Amount = t.Amount.Abs()
	}

	t.OperationDate = time.Now()

	return s.repository.Create(t)
}
