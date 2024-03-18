package simple

type Car struct {
	Type Type
}

func NewCar(car Type) *Car {
	return &Car{
		Type: car,
	}
}

type Owner struct {
	Name Name
}

func NewOwner(name Name) *Owner {
	return &Owner{
		Name: name,
	}
}

type CarOwner struct {
	*Car
	*Owner
}
