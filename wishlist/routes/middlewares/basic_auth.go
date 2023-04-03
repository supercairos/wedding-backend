package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BasicAuth(logger *zap.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == "admin" && password == "password" {
			logger.Info("BasicAuth: OK")
		} else {
			c.Abort()
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}
	}
}
