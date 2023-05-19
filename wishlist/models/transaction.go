package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Transaction struct {
	ID         string  `form:"id" json:"id" binding:"omitempty"`
	Amount     float64 `form:"amount" json:"amount" binding:"required"`
	PaypalTxID *string `form:"paypal_tx_id" json:"paypal_tx_id" binding:"omitempty"`
}

// Transaction is the public data that should hide the SQL implementation
// details from the rest of the code.
type DatabaseTransaction struct {
	ID         string      `db:"id"`
	CreatedAt  time.Time   `db:"created_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
	ItemID     string      `db:"item_id"`
	Amount     float64     `db:"amount"`
	PaypalTxID null.String `db:"paypal_tx_id"`
}

func ToTransactions(i []*DatabaseTransaction) []*Transaction {
	var output []*Transaction
	for _, element := range i {
		output = append(output, ToTransaction(element))
	}

	return output
}

func ToTransaction(i *DatabaseTransaction) *Transaction {
	return &Transaction{
		ID:         i.ID,
		Amount:     i.Amount,
		PaypalTxID: i.PaypalTxID.Ptr(),
	}
}

// TransactionService is the data mapping layer interface, again hiding implementation details.
type TransactionService interface {
	Create(item_id string, item *Transaction) (*DatabaseTransaction, error)
	FindAllByItemID(item_id string) ([]*DatabaseTransaction, error)
}
