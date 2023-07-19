package controller

import (
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"go-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)


type CategoryControllerImpl struct{
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

// Save
func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Read JSON -> read from reader / streaming -> request body

	// # Before
	// decoder := json.NewDecoder(request.Body)

	// categoryCreateRequest := web.CategoryCreateRequest{}
	// err := decoder.Decode(&categoryCreateRequest)
	// helper.PanicIfError(err)

	// # After
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)


	// Conversion
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	// # Before
	// writer.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(writer)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)

	// # After
	helper.WriteToResponseBody(writer, webResponse)
}

// Delete
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Find By Id
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Find All
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
