package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/supercairos/wedding-backend/wishlist/controllers"
	"github.com/supercairos/wedding-backend/wishlist/models/sql"
	"github.com/supercairos/wedding-backend/wishlist/routes/middlewares"
	"go.uber.org/zap"
)

// NewItemRoutes returns a new router for the item resource.
func NewItemRoutes(gin *gin.Engine, logger *zap.Logger, db *sqlx.DB) error {
	is, err := sql.NewItemService(logger, db)
	if err != nil {
		return err
	}

	itemCtrl := controllers.NewItemController(logger, is)

	// GET ITEM
	gin.GET("/items", itemCtrl.GetAll)
	gin.GET("/items/:item_id", itemCtrl.GetByID)
	// CREATE ITEM
	gin.POST("/items", middlewares.BasicAuth(logger), itemCtrl.Post)
	// UPDATE ITEM
	gin.PUT("/items/:item_id", middlewares.BasicAuth(logger), itemCtrl.Put)
	// DELETE ITEM
	gin.DELETE("/items/:item_id", middlewares.BasicAuth(logger), itemCtrl.Delete)

	return nil
}
