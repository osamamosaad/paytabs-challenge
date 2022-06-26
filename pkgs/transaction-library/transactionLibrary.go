package transactionLibrary

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/osamamosaad/paytabs/entities"
	"github.com/osamamosaad/paytabs/pkgs/contracts"
	"github.com/osamamosaad/paytabs/repositories"
	"github.com/osamamosaad/paytabs/utils"
)

type Transaction struct {
	repo       repositories.ITransaction
	accountLib contracts.IAccountLibrary
	mu         sync.Mutex
}

func New(transactionRepo repositories.ITransaction, accountLib contracts.IAccountLibrary) *Transaction {
	return &Transaction{repo: transactionRepo, accountLib: accountLib}
}

func (t *Transaction) Transfer(entity entities.Transaction) (*entities.Transaction, error) {
	err := IsDataValid(entity)
	if err != nil {
		return nil, err
	}

	// Lock the process
	t.mu.Lock()
	defer t.mu.Unlock()

	fromAcc, err := t.accountLib.FindById(entity.From)

	if err != nil {
		return nil, err
	}

	toAcc, err := t.accountLib.FindById(entity.To)
	if err != nil {
		return nil, err
	}

	if fromAcc.Balance < entity.Amount {
		return nil, errors.New(ERR_BALANCE_LESS_THAN_AMOUNT)
	}

	fromAcc.Balance = utils.RoundFloat(fromAcc.Balance-entity.Amount, 2)
	toAcc.Balance = utils.RoundFloat(toAcc.Balance+entity.Amount, 2)

	t.accountLib.Store(fromAcc)
	t.accountLib.Store(toAcc)

	return t.store(entity), nil
}

func (t *Transaction) store(entity entities.Transaction) *entities.Transaction {
	entity.SetID(uuid.New().String())
	entity.TransactionTime = time.Now()
	t.repo.Store(&entity)

	return &entity
}

func IsDataValid(entity entities.Transaction) error {
	if entity.From == "" {
		return errors.New(ERR_MISSING_FROM)
	}

	if entity.To == "" {
		return errors.New(ERR_MISSING_TO)
	}

	if entity.Amount <= 0 {
		return errors.New(ERR_INVALID_AMOUNT)
	}

	if entity.From == entity.To {
		return errors.New(ERR_INVALID_TRANSFER_SAME_ACC)
	}

	return nil
}
