package account

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/pkg/database"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"testing"
)

var conn *gorm.DB

func TestMain(m *testing.M) {
	testDB := database.SetupTestDatabase()
	conn = testDB.Conn
	defer testDB.TearDown()
	os.Exit(m.Run())
}

func TestCreateAccount(t *testing.T) {
	t.Run("create account successfully", func(t *testing.T) {
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		repo := NewRepository(conn, logger)

		acc := &Account{
			Document:             "123456",
			AvailableCreditLimit: decimal.NewFromFloat(132.45),
		}

		service := NewService(repo)
		err := service.Create(context.Background(), acc)

		assert.Nil(t, err)
	})
}
