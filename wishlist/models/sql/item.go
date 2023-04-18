package sql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/supercairos/wedding-backend/wishlist/models"
	"github.com/supercairos/wedding-backend/wishlist/utils"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v4"
)

// ItemService is the implementation of the item data mapping layer
// using SQL.
type ItemService struct {
	conn *sqlx.DB
	log  *zap.Logger
}

// Check it implements the interface
var _ models.ItemService = &ItemService{}

// NewPersonService creates the person service using the given
// connection pool to a postgres DB.
func NewItemService(logger *zap.Logger, conn *sqlx.DB) (*ItemService, error) {
	_, err := conn.Exec(`
CREATE TABLE IF NOT EXISTS items (
	id          SERIAL PRIMARY KEY,
	created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	name	    TEXT NOT NULL,
	paypal_id   TEXT NOT NULL,
	paypal_env  TEXT NOT NULL,
	price	    FLOAT(12, 2) NOT NULL,
	picture_url TEXT
);
`)
	if err != nil {
		return nil, err
	}

	return &ItemService{conn: conn, log: logger}, nil
}

// Create will try to add the person to the DB.
func (s *ItemService) Create(item *models.Item) (*models.DatabaseItem, error) {
	q_insert := `
INSERT INTO items(name, price, paypal_id, paypal_env, picture_url)
VALUES (?,?,?,?,?);`

	result, err := s.conn.Exec(
		q_insert,
		item.Name,
		item.Price,
		item.PaypalId,
		item.PaypalEnv,
		null.StringFromPtr(item.PictureUrl),
	)
	if err != nil {
		return nil, err
	}

	q_select := `
SELECT * FROM items
WHERE id = ?;`

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var output models.DatabaseItem
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

// Update will replace the values of the give person with those provided.
func (s *ItemService) Update(p *models.DatabaseItem) (*models.DatabaseItem, error) {
	if !utils.ValidID(p.ID) {
		return nil, models.ErrNotFound
	}

	q := `
UPDATE items
SET updated_at = NOW(),
    name = ?,
	price = ?,
	paypal_id = ?,
	paypal_env = ?,
	picture_url = ?
WHERE id = ?;`

	result, err := s.conn.Exec(
		q,
		p.Name,
		p.Price,
		p.PaypalId,
		p.PaypalEnv,
		p.PictureUrl,
		p.ID,
	)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	} else if rows == 0 {
		return nil, models.ErrNotFound
	} else if rows > 1 {
		return nil, models.ErrTooManyRows
	}

	q_select := `
SELECT * FROM items
WHERE id = ?;`

	var output models.DatabaseItem
	err = s.conn.Get(
		&output,
		q_select,
		p.ID,
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

func (s *ItemService) FindAll() ([]*models.DatabaseItem, error) {
	q := `
SELECT *, (
	SELECT SUM(transactions.amount) 
	FROM transactions
	WHERE transactions.item_id = items.id
	GROUP BY transactions.item_id
) AS raised 
FROM items;`

	var output []*models.DatabaseItem
	err := s.conn.Select(
		&output,
		q,
	)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// GetByID fetches the person with the given id.
func (s *ItemService) FindByID(id string) (*models.DatabaseItem, error) {
	if !utils.ValidID(id) {
		return nil, models.ErrNotFound
	}

	q := `
SELECT *, (
	SELECT SUM(transactions.amount) 
	FROM transactions
	WHERE transactions.item_id = items.id
	GROUP BY transactions.item_id
) AS raised 
FROM items
WHERE id = ?;`

	var output models.DatabaseItem
	err := s.conn.Get(
		&output,
		q,
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

// Delete removes the item with the given id from the DB.
func (s *ItemService) Delete(id string) error {
	if !utils.ValidID(id) {
		return models.ErrNotFound
	}

	q := `
DELETE FROM items
WHERE id = ?;`

	result, err := s.conn.Exec(
		q,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return models.ErrNotFound
	} else if rows > 1 {
		return models.ErrTooManyRows
	}

	return err
}
