package service

import (
	"context"
	"database/sql"

	// "fmt"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: db,
		Validate: validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.HandleErrorWithPanic(err)

	tx, err := service.DB.Begin()
	helper.HandleErrorWithPanic(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ConvertToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.HandleErrorWithPanic(err)

	tx, err := service.DB.Begin()
	helper.HandleErrorWithPanic(err)
	defer helper.CommitOrRollback(tx)
	
	// temukan id nya terlebih dahulu
	findCategory , err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	// helper.HandleErrorWithPanic(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	category := domain.Category{
		Id: findCategory.Id,
		Name: request.Name,
	}
	// fmt.Println(category)

	// lakukan update
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ConvertToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.HandleErrorWithPanic(err)
	defer helper.CommitOrRollback(tx)

	// temukan id nya terlebih dahulu
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.HandleErrorWithPanic(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category.Id)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorWithPanic(err)
	defer helper.CommitOrRollback(tx)

	// temukan id nya terlebih dahulu
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.HandleErrorWithPanic(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ConvertToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.HandleErrorWithPanic(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ConvertToCategoryResponses(categories)
}
