package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"review_system/database/db_model"
	"time"
)

func New() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("review_system.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(3)

	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("review_system-test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(3)

	return db
}

func DropTestDB() error {
	if err := os.Remove("review_system-test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&db_model.Category{},
		&db_model.Item{},
		&db_model.Review{},
	)
}
