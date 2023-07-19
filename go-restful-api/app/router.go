package app

import (
	"fmt"
	"go-restful-api/controller"
	"go-restful-api/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// # Router After
func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	// Set Router
	// Follow apispec
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Test Router
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello Go RESTful API!")
	})

	router.PanicHandler = exception.ErrorHandler

	return router
}
