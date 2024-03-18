package simple

type CarOwnerService struct {
	*CarService
	*OwnerService
}

func NewCarOwnerService(carService *CarService, ownerService *OwnerService) *CarOwnerService {
	return &CarOwnerService{
		CarService: carService,
		OwnerService: ownerService,
	}
}