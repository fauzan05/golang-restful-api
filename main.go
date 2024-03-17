package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	
	// ketika mengakses route dan ada error maka akan dieksekusi
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.HandleErrorWithPanic(err)
}
