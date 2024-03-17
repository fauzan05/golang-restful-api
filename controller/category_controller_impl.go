package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (categoryController *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := categoryController.CategoryService.Create(r.Context(), categoryCreateRequest)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success create category",
		Data:   categoryResponse,
	}

	helper.ConvertToJson(w, apiResponse)
}

func (categoryController *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)
	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.HandleErrorWithPanic(err)
	categoryUpdateRequest.Id = categoryId
	
	categoryResponse := categoryController.CategoryService.Update(r.Context(), categoryUpdateRequest)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success update category with id : " + strconv.Itoa(categoryResponse.Id),
		Data:   categoryResponse,
	}

	helper.ConvertToJson(w, apiResponse)
}

func (categoryController *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.HandleErrorWithPanic(err)

	categoryController.CategoryService.Delete(r.Context(), categoryId)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success delete category with id : " + strconv.Itoa(categoryId),
		Data:   nil,
	}

	helper.ConvertToJson(w, apiResponse)
}

func (categoryController *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.HandleErrorWithPanic(err)

	getCategory := categoryController.CategoryService.FindById(r.Context(), categoryId)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success show category by id : " + strconv.Itoa(categoryId),
		Data:   getCategory,
	}

	helper.ConvertToJson(w, apiResponse)
}

func (categoryController *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	getCategories := categoryController.CategoryService.FindAll(r.Context())
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success show all category",
		Data:   getCategories,
	}

	helper.ConvertToJson(w, apiResponse)
}
