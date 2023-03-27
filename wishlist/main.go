package main

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/supercairos/wedding-backend/wishlist/routes"
	"github.com/supercairos/wedding-backend/wishlist/utils"
	"go.uber.org/zap"
)

func main() {
	// Use a zap logger
	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		panic(err)
	}

	db, err := utils.NewSqlConnection(logger)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
		return
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/ping", func(c *gin.Context) {
		logger.Info("http ping")
		c.String(http.StatusOK, "pong")
	})

	err = routes.NewItemRoutes(r, logger, db)
	if err != nil {
		logger.Fatal("Failed to create items routes", zap.Error(err))
		return
	}

	err = r.Run(":1337") // listen and serve on 0.0.0.0:1337
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
