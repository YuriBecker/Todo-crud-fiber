package controller

import "gorm.io/gorm"

type AppController struct {
	Todo *todoController
}

func NewAppController(db *gorm.DB) *AppController {
	return &AppController{
		Todo: newTodoController(db),
	}
}
