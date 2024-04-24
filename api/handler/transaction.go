package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/supwr/pismo-transactions/internal/transaction"
	"log/slog"
	"net/http"
)

type TransactionInputDTO struct {
	AccountId       int             `json:"account_id" validate:"required"`
	OperationTypeId int             `json:"operation_type_id" validate:"required"`
	Amount          decimal.Decimal `json:"amount" validate:"required"`
}

type TransactionHandler struct {
	transactionService *transaction.Service
	logger             *slog.Logger
}

func NewTransactionHandler(s *transaction.Service, l *slog.Logger) *TransactionHandler {
	return &TransactionHandler{
		transactionService: s,
		logger:             l,
	}
}

// CreateTransaction godoc
// @Summary      Create transaction
// @Description  Add new transaction
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request   body      TransactionInputDTO  true  "Transaction properties"
// @Success      201
// @Failure      500
// @Failure      400
// @Router       /transactions [post]
func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var err error
	var input TransactionInputDTO

	if err = ctx.BindJSON(&input); err != nil {
		h.logger.ErrorContext(ctx, "error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	validation := validate(input).Errors
	if len(validation) > 0 {
		h.logger.ErrorContext(ctx, "invalid payload", slog.Any("error", err))
		ctx.JSON(http.StatusBadRequest, validation)
		return
	}

	transact := &transaction.Transaction{
		AccountID:       input.AccountId,
		OperationTypeID: input.OperationTypeId,
		Amount:          input.Amount,
	}

	if err = h.transactionService.Create(ctx, transact); err != nil {
		h.logger.ErrorContext(ctx, "error creating transaction", slog.Any("error", err))
		if errors.Is(err, transaction.ErrOperationTypeNotFound) || errors.Is(err, transaction.ErrAccountNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrCreateTransaction.Error(),
		})
		return
	}

	h.logger.InfoContext(ctx, "transaction created successfully", slog.Any("transaction", transact))
	ctx.JSON(http.StatusCreated, nil)
	return
}
