package models

import (
	"time"

	"github.com/supercairos/wedding-backend/wishlist/utils"
	"gopkg.in/guregu/null.v4"
)

type Item struct {
	ID         string  `form:"id" json:"id" binding:"omitempty"`
	Name       string  `form:"name" json:"name" binding:"required"`
	Price      float64 `form:"price" json:"price" binding:"required"`
	PictureUrl *string `form:"picture_url" json:"picture_url" binding:"omitempty"`
	Raised     float64 `form:"raised" json:"raised" binding:"omitempty"`
}

// Item is the public data that should hide the SQL implementation
// details from the rest of the code.
type DatabaseItem struct {
	ID         string      `db:"id"`
	CreatedAt  time.Time   `db:"created_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
	Name       string      `db:"name"`
	Price      float64     `db:"price"`
	PictureUrl null.String `db:"picture_url"`
	Raised     null.Float  `db:"raised"`
}

func ToItems(i []*DatabaseItem) []*Item {
	var output []*Item
	for _, element := range i {
		output = append(output, ToItem(element))
	}

	return output
}

func ToItem(i *DatabaseItem) *Item {
	return &Item{
		ID:         i.ID,
		Name:       i.Name,
		Price:      i.Price,
		PictureUrl: i.PictureUrl.Ptr(),
		Raised:     utils.ToFloatValue(i.Raised),
	}
}

// ItemService is the data mapping layer interface, again hiding implementation details.
type ItemService interface {
	Create(item *Item) (*DatabaseItem, error)
	FindAll() ([]*DatabaseItem, error)
	FindByID(id string) (*DatabaseItem, error)
	Update(item *DatabaseItem) (*DatabaseItem, error)
	Delete(id string) error
}
