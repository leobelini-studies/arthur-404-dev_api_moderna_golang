package handler

import (
	"github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
