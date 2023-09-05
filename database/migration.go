package database

import "github.com/yuribecker/todo-crud-fiber/model"

func Migrate() error {
	err := DB.AutoMigrate(&model.Todo{})

	return err
}
