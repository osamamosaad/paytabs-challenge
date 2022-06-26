package contracts

import "github.com/osamamosaad/paytabs/entities"

// ITransactionLibrary interface transaction lirary
type ITransactionLibrary interface {
	Transfer(entity entities.Transaction) (*entities.Transaction, error)
}
