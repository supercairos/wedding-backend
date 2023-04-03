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
func NewTransactionRoutes(gin *gin.Engine, logger *zap.Logger, db *sqlx.DB) error {
	ts, err := sql.NewTransactionService(logger, db)
	if err != nil {
		return err
	}

	transactionCtrl := controllers.NewTransactionController(ts)

	// GET TRANSACTION
	gin.GET("/items/:item_id/transactions", middlewares.BasicAuth(logger), transactionCtrl.GetAll(logger))
	// POST TRANSACTION
	gin.POST("/items/:item_id/transactions", transactionCtrl.Post(logger))

	return nil
}
