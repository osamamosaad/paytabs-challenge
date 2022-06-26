package entities

type Accounts []*Account

type Account struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance,string"`
}

func (a *Account) SetID(ID string) {
	a.ID = ID
}

func (a Account) GetID() string {
	return a.ID
}
