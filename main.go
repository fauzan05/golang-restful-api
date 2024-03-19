package main

import (
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:8000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.HandleErrorWithPanic(err)
	// go run main.go wire_gen.go
}
