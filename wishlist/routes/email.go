package routes

import (
	"github.com/gin-gonic/gin"
	siblib "github.com/sendinblue/APIv3-go-library/v2/lib"
	"github.com/supercairos/wedding-backend/wishlist/controllers"
	"github.com/supercairos/wedding-backend/wishlist/models/sendinblue"
	"go.uber.org/zap"
)

// NewItemRoutes returns a new router for the item resource.
func NewEmailRoute(gin *gin.Engine, logger *zap.Logger, sib *siblib.APIClient) error {
	es, err := sendinblue.NewEmailService(logger, sib)
	if err != nil {
		return err
	}

	emailCtrl := controllers.NewEmailController(es)

	// SEND EMAIL
	gin.POST("/email", emailCtrl.Send(logger))

	return nil
}
