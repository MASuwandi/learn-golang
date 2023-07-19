package main

import (
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Init DB
	db := app.NewDB()
	validate := validator.New()

	// Init
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// # Router Before
	// router := httprouter.New()

	// // Set Router
	// // Follow apispec
	// router.GET("/api/categories", categoryController.FindAll)
	// router.GET("/api/categories/:categoryId", categoryController.FindById)
	// router.POST("/api/categories", categoryController.Create)
	// router.PUT("/api/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Test Router
	// router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)  {
	// 	fmt.Fprint(writer, "Hello Go RESTful API!")
	// })

	// Panic Handler / set custom error handler
	// router.PanicHandler = exception.ErrorHandler

	// # After Router
	router := app.NewRouter(categoryController)

	// Follow apispec
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
