package transaction

import (
	"github.com/supwr/pismo-transactions/internal/entity"
	"github.com/supwr/pismo-transactions/internal/usecase/account"
	"github.com/supwr/pismo-transactions/internal/usecase/operation_type"
	"github.com/supwr/pismo-transactions/pkg/clock"
	"slices"
)

type Service struct {
	repository           RepositoryInterface
	operationTypeService *operation_type.Service
	accountService       *account.Service
	clock                clock.Clock
}

func NewService(r RepositoryInterface, o *operation_type.Service, a *account.Service, c clock.Clock) *Service {
	return &Service{repository: r, operationTypeService: o, accountService: a, clock: c}
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

	t.OperationDate = s.clock.Now()

	return s.repository.Create(t)
}
