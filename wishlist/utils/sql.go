package utils

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// NewSQL creates and SQL connection using environment variables
// to configure.
func NewSqlConnection(logger *zap.Logger) (*sqlx.DB, error) {
	host := strings.TrimSpace(os.Getenv("DATABASE_HOST"))
	port := strings.TrimSpace(os.Getenv("DATABASE_PORT"))
	user := strings.TrimSpace(os.Getenv("DATABASE_USER"))
	password := strings.TrimSpace(os.Getenv("DATABASE_PASSWORD"))
	db := strings.TrimSpace(os.Getenv("DATABASE_DBNAME"))

	info := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		db,
	)

	logger.Info("connecting to database", zap.String("info", info))
	conn, err := sqlx.Connect("mysql", info)
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		logger.Error("failed to ping database", zap.Error(err))
		return nil, err
	}

	return conn, nil
}
