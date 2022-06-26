package entities

import "time"

type Transaction struct {
	ID              string    `json:"id,omitempty"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
}

func (t *Transaction) SetID(ID string) {
	t.ID = ID
}

func (t Transaction) GetID() string {
	return t.ID
}
