package middleware

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-KEY") == "12345" {
		// BERHASIL
		middleware.Handler.ServeHTTP(w, r)
	} else {
		// GAGAL
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		apiResponse := web.ApiResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   nil,
		}
		helper.ConvertToJson(w, apiResponse)	
	}
}