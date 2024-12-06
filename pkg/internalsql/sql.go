package internalsql

import (
	"database/sql"
	"fmt"
	"social-network/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(config *config.ConfigTypes) (*sql.DB, error) {
	sourceDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE,
	)

	db, err := sql.Open("mysql", sourceDB)

	return db, err
}
