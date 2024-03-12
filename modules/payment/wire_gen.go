// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package invoice

import (
	"neema.co.za/rest/modules/payment/internal/api"
	"neema.co.za/rest/modules/payment/internal/repository"
	"neema.co.za/rest/modules/payment/internal/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
	"neema.co.za/rest/utils/managers"
)

// Injectors from wire.go:

func BuildApi(dependencyManager *managers.DependencyManager) *api.Api {
	databaseDatabase := database.GetDatabase()
	repositoryRepository := repository.NewRepository(databaseDatabase)
	serviceService := service.NewService(repositoryRepository, dependencyManager)
	fiberApp := app.NewFiberApp()
	apiApi := api.NewApi(serviceService, fiberApp)
	return apiApi
}