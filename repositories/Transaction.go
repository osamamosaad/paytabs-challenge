package repositories

import (
	"github.com/osamamosaad/paytabs/entities"
	"github.com/osamamosaad/paytabs/storage"
)

type ITransaction interface {
	Store(entity *entities.Transaction)
}

type Transaction struct {
	storage *storage.Storage
}

func NewTransaction(storage *storage.Storage) ITransaction {
	return &Transaction{storage.SetTableName("Transaction")}
}

func (t Transaction) Store(entity *entities.Transaction) {
	t.storage.Store(entity)
}
