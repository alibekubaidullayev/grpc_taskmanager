package core

import (
	"fmt"
	"sync"
	"tm/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

func initDB() {
	cfg := GetConfig()
	dbLink := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DbHost, cfg.DbUser, cfg.DbPasw, cfg.DbName, cfg.DbPort)

	db, err = gorm.Open(postgres.Open(dbLink), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Task{})
}

func GetDB() (*gorm.DB, error) {
	once.Do(initDB)
	return db, err
}
