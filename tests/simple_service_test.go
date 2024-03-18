package tests

import (
	"fmt"
	"golang-restful-api/simple"
	"testing"
)

func TestDependencyInjection(t *testing.T) {
	categoryService, err := simple.InitializedService(false, "Hai")
	if err != nil {
		fmt.Println("Errornya : ", err)
	} else {
		fmt.Println(categoryService.SimpleRepository)
	}
}

func TestDependencyInjectionCarOwner(t *testing.T) {
	carOwner := simple.InitializedCarOwner("Ferrari", "Fauzan")
	fmt.Println(carOwner.Type)
	fmt.Println(carOwner.Name)
}

func TestDependencyInjectionConnection(t *testing.T) {
	connection, function := simple.InitializedConnection("Jumanji")
	defer function()
	fmt.Println("Nama file : ", connection.File.Name)
}
