package transactionLibrary

import (
	"github.com/osamamosaad/paytabs/entities"
	"github.com/osamamosaad/paytabs/repositories"
	"github.com/osamamosaad/paytabs/storage"
)

func getMockAccountRepo(s *storage.Storage, isStoreEnable bool) repositories.IAccount {
	return &mockAccountRepo{
		storage:       s.SetTableName("Account"),
		accountRepo:   repositories.NewAccount(s),
		isStoreEnable: isStoreEnable,
	}
}

type mockAccountRepo struct {
	storage       *storage.Storage
	accountRepo   repositories.IAccount
	isStoreEnable bool
}

func (m *mockAccountRepo) List() (entities.Accounts, error) {
	return entities.Accounts{}, nil
}

func (m *mockAccountRepo) FindById(ID string) (*entities.Account, error) {
	return m.accountRepo.FindById(ID)
}

func (m *mockAccountRepo) Store(entity entities.EntityInterface) (interface{}, error) {
	if m.isStoreEnable {
		return m.accountRepo.Store(entity)
	}

	return entity, nil
}

type mockTransactionRepo struct {
}

func (t mockTransactionRepo) Store(entity *entities.Transaction) {
	// Do noting
}
