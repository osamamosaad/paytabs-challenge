package accountLibrary

import (
	"github.com/osamamosaad/paytabs/entities"
	"github.com/osamamosaad/paytabs/pkgs/contracts"
	"github.com/osamamosaad/paytabs/repositories"
)

type account struct {
	repo repositories.IAccount
}

func New(AccountRepo repositories.IAccount) contracts.IAccountLibrary {
	return &account{AccountRepo}
}

func (a account) Store(entity *entities.Account) {
	a.repo.Store(entity)
}

func (a account) FindById(ID string) (*entities.Account, error) {
	return a.repo.FindById(ID)
}

func (a account) FindAll() (entities.Accounts, error) {
	return a.repo.List()
}
