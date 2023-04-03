package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

	sib, err := utils.NewSibClient(logger)
	if err != nil {
		logger.Fatal("Failed to connect to sendinblue", zap.Error(err))
		return
	}

	r := gin.New()
	r.SetTrustedProxies(nil)

	r.Use(cors.Default())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/ping", func(c *gin.Context) {
		logger.Info("http ping")
		c.String(http.StatusOK, `{"message": "pong"}`)
	})

	err = routes.NewItemRoutes(r, logger, db)
	if err != nil {
		logger.Fatal("Failed to create items routes", zap.Error(err))
		return
	}

	err = routes.NewTransactionRoutes(r, logger, db)
	if err != nil {
		logger.Fatal("Failed to create transactions routes", zap.Error(err))
		return
	}

	err = routes.NewEmailRoute(r, logger, sib)
	if err != nil {
		logger.Fatal("Failed to create email routes", zap.Error(err))
		return
	}

	port := utils.GetEnv("PORT", "1337")
	err = r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:1337
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}

	logger.Info("Server started", zap.String("port", port))
}
