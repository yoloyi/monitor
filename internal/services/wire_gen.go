// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"go-admin/internal/models"
	"go-admin/internal/models/repositories"
	"go-admin/internal/services/auth"
	auth2 "go-admin/internal/utils/auth"
)

// Injectors from wire.go:

func NewAuthService() *auth.Service {
	db := models.GetDB()
	user := repositories.NewUser(db)
	authAuth := auth2.NewAuth()
	auth3 := auth.NewAuth(user, authAuth)
	return auth3
}
