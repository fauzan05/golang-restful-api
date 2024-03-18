//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"os"
	"io"
)

func InitializedService(isError bool, word string)(*SimpleService, error){
	 wire.Build(
		NewSimpleRepository, NewSimpleService,
	 )
	 return nil, nil
}

func InitializedDatabase() *DatabaseRepository {
	wire.Build(
		NewDatabaseMySQL, NewDatabaseMongoDB, NewDatabaseRepository,
	)
	return nil
}

var carSet = wire.NewSet(NewCarRepository, NewCarService)
var ownerSet = wire.NewSet(NewOwnerRepository, NewOwnerService)

func InitializedCarOwnerService(Type Type, Speed Speed, Name Name, Age Age) *CarOwnerService {
	wire.Build(
		carSet, ownerSet, NewCarOwnerService,
	)
	return nil
}

/*
	jika ada provider yang terikat dengan interface,
	maka wajib memberi tahu bahwa provider tersebut
	memiliki interface

*/

var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
	// (interfacenya, implementasinya)
)

func InitializedHelloService(Age string) *HelloService {
	wire.Build(
		// NewSayHelloImpl, NewSayHelloService,
		HelloSet, NewHelloService,
	)
	return nil
}


var CarOwnerSet = wire.NewSet(
	NewCar,
	NewOwner,
)

func InitializedCarOwner(car Type, name Name) *CarOwner {
	wire.Build(
		CarOwnerSet,
		wire.Struct(new(CarOwner), "Car", "Owner"),
		// parameter bagian "Car" dan "Owner adalah field mana yang akan diinject pada struct CarOwner. jika malas menyebutkan satu satu, bisa menggunakan operator * agar semua field di inject"
	)
	return nil
}


var carValue = &Car{}
var ownerValue = &Owner{}

// langsung menggunakan struct-nya tanpa harus membuat providernya
func InitializedCarOwnerUsingValue() *CarOwner {
	wire.Build(wire.Value(carValue), wire.Value(ownerValue), wire.Struct(new(CarOwner), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// cara membuat provider Configuration secara otomatis
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}


func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}