package main

import (
	"golearn-restful-api/controller"
	"golearn-restful-api/repository"
	"golearn-restful-api/service"

	"github.com/julienschmidt/httprouter"
)

func main() {

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
}
