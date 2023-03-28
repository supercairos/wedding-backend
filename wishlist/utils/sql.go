package utils

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// NewSQL creates and SQL connection using environment variables
// to configure.
func NewSqlConnection(logger *zap.Logger) (*sqlx.DB, error) {
	dsn := strings.TrimSpace(os.Getenv("DSN"))
	logger.Info("connecting to database", zap.String("info", dsn))
	conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		logger.Error("failed to ping database", zap.Error(err))
		return nil, err
	}
	logger.Info("connected to database")
	return conn, nil
}
