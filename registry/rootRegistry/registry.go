package rootRegistry

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/registry/authRegistry"
	"awesomeProject/registry/tokensRegistry"
	"awesomeProject/registry/userRegistry"
	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRegistry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) IRegistry {
	return &Registry{db}
}

func (r *Registry) NewAppController() controller.AppController {
	return controller.AppController{
		User:   userRegistry.NewRegistry(r.db),
		Auth:   authRegistry.NewRegistry(r.db),
		Tokens: tokensRegistry.NewRegistry(r.db),
	}
}