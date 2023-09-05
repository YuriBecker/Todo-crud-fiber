package database

import (
	"fmt"

	"github.com/yuribecker/todo-crud-fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (err error) {
	cfg := config.Get().DB
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn, cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
