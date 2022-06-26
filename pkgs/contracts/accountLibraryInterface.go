package contracts

import "github.com/osamamosaad/paytabs/entities"

// IAccountLibrary interface account lirary
type IAccountLibrary interface {
	Store(entity *entities.Account)
	FindById(ID string) (*entities.Account, error)
	FindAll() (entities.Accounts, error)
}
