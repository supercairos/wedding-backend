package sql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/supercairos/wedding-backend/wishlist/models"
	"github.com/supercairos/wedding-backend/wishlist/utils"
)

// TransactionService is the implementation of the item data mapping layer
// using SQL.
type TransactionService struct {
	conn *sqlx.DB
}

// Check it implements the interface
var _ models.TransactionService = &TransactionService{}

// NewPersonService creates the person service using the given
// connection pool to a postgres DB.
func NewTransactionService(conn *sqlx.DB) (*TransactionService, error) {
	_, err := conn.Exec(`
CREATE TABLE IF NOT EXISTS transactions (
	id           SERIAL PRIMARY KEY,
	created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at   TIMESTAMP NOT NULL DEFAULT NOW(),
	item_id	     TEXT NOT NULL,
	amount	     FLOAT(12, 2) NOT NULL,
	paypal_tx_id TEXT NOT NULL
);
`)
	if err != nil {
		return nil, err
	}

	return &TransactionService{conn: conn}, nil
}

// Create will try to add the person to the DB.
func (s *TransactionService) Create(item_id string, transaction *models.Transaction) (*models.DatabaseTransaction, error) {
	if !utils.ValidID(item_id) {
		return nil, models.ErrNotFound
	}

	// Assert that item exist
	q_item := `
SELECT * FROM items
WHERE id = ?;`

	var item models.DatabaseItem
	err := s.conn.Get(
		&item,
		q_item,
		item_id,
	)

	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// Insert transactions
	q_insert := `
INSERT INTO transactions(item_id, amount, paypal_tx_id)
VALUES (?,?,?);`

	result, err := s.conn.Exec(
		q_insert,
		item.ID,
		transaction.Amount,
		transaction.PaypalTxID,
	)
	if err != nil {
		return nil, err
	}

	// Query last transaction
	q_select := `
SELECT * FROM transactions
WHERE id = ?;`

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var output models.DatabaseTransaction
	err = s.conn.Get(
		&output,
		q_select,
		id,
	)

	// Replace the SQL error with our own error type.
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	} else {
		return &output, nil
	}
}

func (s *TransactionService) FindAllByItemID(item_id string) ([]*models.DatabaseTransaction, error) {
	if !utils.ValidID(item_id) {
		return nil, models.ErrNotFound
	}

	// Assert that item exist
	q_item := `
SELECT * FROM items
WHERE id = ?;`

	var item models.DatabaseItem
	err := s.conn.Get(
		&item,
		q_item,
		item_id,
	)

	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	q_select := `
SELECT * FROM transactions
WHERE item_id = ?;`

	var transactions []*models.DatabaseTransaction
	err = s.conn.Select(
		&transactions,
		q_select,
		item.ID,
	)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
