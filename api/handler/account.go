package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/supwr/pismo-transactions/entity"
	"github.com/supwr/pismo-transactions/usecase/account"
	"log/slog"
	"net/http"
	"strconv"
)

type AccountInputDTO struct {
	DocumentNumber entity.Document `json:"document_number" swaggertype:"string" validate:"required"`
}

type AccountHandler struct {
	AccountService *account.Service
	logger         *slog.Logger
}

func NewAccountHandler(s *account.Service, l *slog.Logger) *AccountHandler {
	return &AccountHandler{
		AccountService: s,
		logger:         l,
	}
}

// CreateAccount godoc
// @Summary      Create account
// @Description  Add new account
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        request   body      AccountInputDTO  true  "Account properties"
// @Success      201
// @Failure      500
// @Failure      400
// @Router       /accounts [post]
func (h *AccountHandler) CreateAccount(ctx *gin.Context) {
	var err error
	var input AccountInputDTO

	if err = ctx.BindJSON(&input); err != nil {
		h.logger.Error("error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	validation := validate(input).Errors
	if len(validation) > 0 {
		h.logger.Error("invalid payload", slog.Any("error", err))
		ctx.JSON(http.StatusBadRequest, validation)
		return
	}

	acc := &entity.Account{Document: input.DocumentNumber}

	if err = h.AccountService.Create(acc); err != nil {
		h.logger.Error("error creating account", slog.Any("error", err))
		if errors.Is(err, account.ErrAccountAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrCreateAccount.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
	return
}

// GetAccountById godoc
// @Summary      Show account details
// @Description  Get account by id
// @Tags         Accounts
// @Produce      json
// @Param        accountId   path      integer  true  "Account id"
// @Success      200
// @Failure      500
// @Failure      404
// @Router       /accounts/{accountId} [get]
func (h *AccountHandler) GetAccountById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("accountId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	acc, err := h.AccountService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if acc == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, acc)
}
