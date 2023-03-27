package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/supercairos/wedding-backend/wishlist/controllers"
	"github.com/supercairos/wedding-backend/wishlist/models/sql"
	"go.uber.org/zap"
)

// NewItemRoutes returns a new router for the item resource.
func NewItemRoutes(gin *gin.Engine, logger *zap.Logger, db *sqlx.DB) error {
	is, err := sql.NewItemService(db)
	if err != nil {
		return err
	}

	ts, err := sql.NewTransactionService(db)
	if err != nil {
		return err
	}

	itemCtrl := controllers.NewItemController(is, ts)
	transactionCtrl := controllers.NewTransactionController(ts)

	group := gin.Group("/items")
	// GET ITEM
	group.GET("/", itemCtrl.GetAll(logger))
	group.GET("/:item_id", itemCtrl.GetByID(logger))
	// CREATE ITEM
	group.POST("/", itemCtrl.Post(logger))
	// UPDATE ITEM
	group.PUT("/:item_id", itemCtrl.Put(logger))
	// DELETE ITEM
	group.DELETE("/:item_id", itemCtrl.Delete(logger))

	// GET TRANSACTION
	group.GET("/:item_id/transactions", transactionCtrl.GetAll(logger))
	// POST TRANSACTION
	group.POST("/:item_id/transactions", transactionCtrl.Post(logger))

	return nil
}
