package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/supercairos/wedding-backend/wishlist/models"
	"go.uber.org/zap"
)

// Item provides the handlers for the Item entity.
type EmailController struct {
	EmailService models.EmailService
}

// NewItem creates the controller using the given data mapper for
// Items.
func NewEmailController(es models.EmailService) *EmailController {
	return &EmailController{
		EmailService: es,
	}
}

func (p *EmailController) Send(logger *zap.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data models.Email
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

		if err := p.EmailService.Send(&data); err != nil {
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
			gin.H{"message": "email sent"},
		)
	}
}
