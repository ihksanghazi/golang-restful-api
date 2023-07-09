package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	validate := validator.New()
	DB:= app.NewDB()
	categoryRepository:= repository.NewCategoryRepository()
	categoryService:=service.NewCategoryService(categoryRepository, DB, validate)
	categoryController:=controller.NewCategoryController(categoryService)
	router:= app.Router(categoryController)

	server:= http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err:=server.ListenAndServe()
	helper.PanicIfError(err)
}