package account

import "github.com/supwr/pismo-transactions/entity"

type RepositoryInterface interface {
	Create(account *entity.Account) error
	FindById(id int) (*entity.Account, error)
	FindByDocument(document entity.Document) (*entity.Account, error)
}
