package main

import (
	"go-dependency-injection/app"
	"go-dependency-injection/controller"
	"go-dependency-injection/helper"
	"go-dependency-injection/middleware"
	"go-dependency-injection/repository"
	"go-dependency-injection/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8088",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
