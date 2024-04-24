package transaction

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/supwr/pismo-transactions/internal/account"
	"github.com/supwr/pismo-transactions/pkg/clock"
	"slices"
)

type Service struct {
	repository     RepositoryInterface
	accountService *account.Service
	clock          clock.Clock
}

func NewService(r RepositoryInterface, a *account.Service, c clock.Clock) *Service {
	return &Service{repository: r, accountService: a, clock: c}
}

func (s *Service) Create(ctx context.Context, t *Transaction) error {
	var negAmountTransactions = []int{OperationTypeCashBuy, OperationTypeInstallmentBuy, OperationTypeWithdraw}

	acc, err := s.accountService.FindById(ctx, t.AccountID)
	if err != nil {
		return err
	}

	if acc == nil {
		return ErrAccountNotFound
	}

	if _, exists := Operations[t.OperationTypeID]; !exists {
		return ErrOperationTypeNotFound
	}

	if slices.Contains(negAmountTransactions, t.OperationTypeID) {
		t.Amount = t.Amount.Abs().Neg()

		if acc.AvailableCreditLimit.Add(t.Amount).LessThan(decimal.Zero) {
			return ErrInsuficientFunds
		}
	} else {
		t.Amount = t.Amount.Abs()
	}

	t.OperationDate = s.clock.Now()

	acc.AvailableCreditLimit = acc.AvailableCreditLimit.Add(t.Amount)

	if err = s.accountService.UpdateCreditLimit(ctx, acc); err != nil {
		return err
	}

	return s.repository.Create(ctx, t)
}
