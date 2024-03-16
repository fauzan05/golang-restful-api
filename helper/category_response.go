package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
)

func ConvertToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ConvertToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	// mengonversi ke bentuk category response
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ConvertToCategoryResponse(category))
	}
	return categoryResponses
}