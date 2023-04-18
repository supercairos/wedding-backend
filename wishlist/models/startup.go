package models

import (
	"time"
)

// Startup is the public data that should hide the SQL implementation
// details from the rest of the code.
type StartupItem struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Date      string    `db:"date"`
}

// ItemService is the data mapping layer interface, again hiding implementation details.
type StartupService interface {
	Create() error
}
