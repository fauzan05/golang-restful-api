package tests

import (
	"database/sql"
	"encoding/json"

	// "fmt"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

func setupTest() http.Handler {
	db := app.NewDBTest()
	truncateCategory(db)
	router := setupRouter(db)
	return router
}

var localhost string = "http://localhost:8000/api"

func TestCreateCategorySuccess(t *testing.T) {
	router := setupTest()
	requestBody := strings.NewReader(`{"name" : "Mainan"}`)
	request := httptest.NewRequest(http.MethodPost, localhost+"/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "12345")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	helper.HandleErrorWithPanic(err)
	// fmt.Println(string(body))
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	// fmt.Println(responseBody["data"])
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Success create category", responseBody["status"])
	assert.Equal(t, "Mainan", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	router := setupTest()
	// buat datanya
	requestBody := strings.NewReader(`{"name" : "Mainan"}`)
	request := httptest.NewRequest(http.MethodPost, localhost+"/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "12345")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// update datanya
	requestBody = strings.NewReader(`{"name" : "Makanan"}`)
	request = httptest.NewRequest(http.MethodPut, localhost+"/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "12345")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	helper.HandleErrorWithPanic(err)
	// fmt.Println(string(body))
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	// fmt.Println(responseBody["data"])
	assert.Equal(t, 200, response.StatusCode)
}
