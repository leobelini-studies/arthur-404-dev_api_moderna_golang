package config

import (
	"os"

	"github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Sqlite DB not found, creating: %s", dbPath)

		// Create DB file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Sqlite directory creation error: %s", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Sqlite file creation error: %s", err)
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Sqlite opening error: %s", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Sqlite migration error: %s", err)
		return nil, err
	}

	// Return DB
	return db, nil
}
