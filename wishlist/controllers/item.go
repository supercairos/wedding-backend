package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/supercairos/wedding-backend/wishlist/models"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v4"
)

// Item provides the handlers for the Item entity.
type ItemController struct {
	ItemService models.ItemService
	Log         *zap.Logger
}

// NewItem creates the controller using the given data mapper for
// Items.
func NewItemController(logger *zap.Logger, is models.ItemService) *ItemController {
	return &ItemController{
		ItemService: is,
		Log:         logger,
	}
}

// Post will create a new Item from the given data, if the form is valid.
func (p *ItemController) Post(c *gin.Context) {
	var data models.Item
	if err := c.ShouldBindWith(&data, binding.JSON); err != nil {
		c.Error(err)
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": "invalid data.",
				"form":    data,
				"error":   err.Error(),
			},
		)
		c.Abort()
		return
	}

	item, err := p.ItemService.Create(&data)
	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error"},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusCreated,
		models.ToItem(item),
	)
}

// Put will perform an update of a Item.
func (p *ItemController) Put(c *gin.Context) {
	var data models.Item
	if err := c.ShouldBindWith(&data, binding.JSON); err != nil {
		c.Error(err)
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": "invalid data.",
				"form":    data,
				"error":   err.Error(),
			},
		)
		c.Abort()
		return
	}

	item, err := p.ItemService.FindByID(c.Param("item_id"))
	if err == models.ErrNotFound {
		c.Error(err)
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "item was not found"},
		)
		c.Abort()
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	item, err = p.ItemService.Update(&models.DatabaseItem{
		ID:         item.ID,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
		Name:       data.Name,
		Price:      data.Price,
		PictureUrl: null.StringFromPtr(data.PictureUrl),
	})
	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		models.ToItem(item),
	)
}

// Get will fetch all Items.
func (p *ItemController) GetAll(c *gin.Context) {
	dbitems, err := p.ItemService.FindAll()
	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	if len(dbitems) == 0 {
		c.JSON(
			http.StatusOK,
			[]models.DatabaseItem{},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		models.ToItems(dbitems),
	)
}

// Get will fetch a Item by ID.
func (p *ItemController) GetByID(c *gin.Context) {
	dbitem, err := p.ItemService.FindByID(c.Param("item_id"))
	if err == models.ErrNotFound {
		c.Error(err)
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "item was not found"},
		)
		c.Abort()
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		models.ToItem(dbitem),
	)
}

// Delete will remove a Item from the DB.
func (p *ItemController) Delete(c *gin.Context) {
	err := p.ItemService.Delete(c.Param("item_id"))
	if err == models.ErrNotFound {
		c.Error(err)
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "item was not found"},
		)
		c.Abort()
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "deleted"},
	)
}
