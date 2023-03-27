package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/supercairos/wedding-backend/wishlist/models"
	"go.uber.org/zap"
)

type TransactionController struct {
	TransactionService models.TransactionService
}

func NewTransactionController(ts models.TransactionService) *TransactionController {
	return &TransactionController{
		TransactionService: ts,
	}
}

// Post will create a new Item from the given data, if the form is valid.
func (p *TransactionController) Post(logger *zap.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data models.Transaction
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

		transaction, err := p.TransactionService.Create(c.Param("item_id"), &data)
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
				gin.H{"message": "internal error"},
			)
			c.Abort()
			return
		}

		c.JSON(
			http.StatusCreated,
			models.ToTransaction(transaction),
		)
	}
}

// Get will fetch all Items.
func (p *TransactionController) GetAll(logger *zap.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		transactions, err := p.TransactionService.FindAllByItemID(c.Param("item_id"))
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

		if len(transactions) == 0 {
			c.JSON(
				http.StatusOK,
				[]models.DatabaseItem{},
			)
			c.Abort()
			return
		}

		c.JSON(
			http.StatusOK,
			models.ToTransactions(transactions),
		)
	}
}
