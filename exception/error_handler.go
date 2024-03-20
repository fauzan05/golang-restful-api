package exception

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if notFoundError(w, err) {
		return
	}

	if validationError(w, err) {
		return
	}

	internalServerError(w, err)
}

func validationError(w http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		apiResponse := web.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(), // mengubah errornya menjadi string
		}
		helper.ConvertToJson(w, apiResponse)
		helper.SaveToLogError(exception.Error())
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		
		apiResponse := web.ApiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helper.ConvertToJson(w, apiResponse)
		helper.SaveToLogError(exception)
		return true
		} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	apiResponse := web.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	helper.SaveToLogError(err)
	helper.ConvertToJson(w, apiResponse)
}
