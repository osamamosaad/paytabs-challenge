package repositories

import (
	"github.com/google/uuid"
	"github.com/osamamosaad/paytabs/entities"
	"github.com/osamamosaad/paytabs/storage"
)

type IAccount interface {
	List() (entities.Accounts, error)
	FindById(ID string) (*entities.Account, error)
	Store(entity entities.EntityInterface) (interface{}, error)
}

type Account struct {
	Storage *storage.Storage
}

func NewAccount(storage *storage.Storage) IAccount {

	return &Account{storage.SetTableName("Account")}
}

func (l *Account) List() (entities.Accounts, error) {
	var accountList entities.Accounts
	results, err := l.Storage.FindAll()
	if err != nil {
		return nil, err
	}

	for _, account := range results {
		accountList = append(accountList, account.(*entities.Account))
	}

	return accountList, nil
}

func (l *Account) FindById(ID string) (*entities.Account, error) {
	result, err := l.Storage.FindById(ID)

	if err != nil {
		return nil, err
	}

	return result.(*entities.Account), err
}

func (l *Account) Store(entity entities.EntityInterface) (interface{}, error) {
	if entity.GetID() == "" {
		entity.SetID(uuid.New().String())
	}

	return l.Storage.FindById(entity.GetID())
}
