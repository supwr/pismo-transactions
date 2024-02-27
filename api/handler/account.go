package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/pismo-transactions/usecase/account"
	"log/slog"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	AccountService *account.Service
	logger         *slog.Logger
}

func NewAccountHandler(service *account.Service, logger *slog.Logger) *AccountHandler {
	return &AccountHandler{
		AccountService: service,
		logger:         logger,
	}
}

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
			"error": err,
		})
		return
	}

	if acc == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, acc)
}
