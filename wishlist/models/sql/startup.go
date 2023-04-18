package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/supercairos/wedding-backend/wishlist/models"
	"go.uber.org/zap"
)

// StartupService is the implementation of the item data mapping layer
// using SQL.
type StartupService struct {
	conn *sqlx.DB
	log  *zap.Logger
}

// Check it implements the interface
var _ models.StartupService = &StartupService{}

// NewPersonService creates the person service using the given
// connection pool to a postgres DB.
func NewStartupService(logger *zap.Logger, conn *sqlx.DB) (*StartupService, error) {
	_, err := conn.Exec(`
CREATE TABLE IF NOT EXISTS startup (
	id          SERIAL PRIMARY KEY,
	created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	date	    TIMESTAMP NOT NULL DEFAULT NOW()
);
`)
	if err != nil {
		return nil, err
	}

	return &StartupService{conn: conn, log: logger}, nil
}

// Create will try to add the person to the DB.
func (s *StartupService) Create() error {
	q_insert := `
INSERT INTO startup() VALUES();`

	_, err := s.conn.Exec(
		q_insert,
	)
	if err != nil {
		return err
	}

	return nil
}
