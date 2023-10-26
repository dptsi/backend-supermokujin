// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-dependency-injection/app"
	"go-dependency-injection/controller"
	"go-dependency-injection/middleware"
	"go-dependency-injection/repository"
	"go-dependency-injection/service"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer(option ...validator.Option) *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New(option...)
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	router := app.NewRouter(categoryControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)), service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))