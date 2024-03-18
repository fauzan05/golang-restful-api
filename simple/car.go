package simple

type Type string
type Speed int

type CarRepository struct {
	Type Type
	Speed Speed
}

func NewCarRepository(Type Type, Speed Speed) *CarRepository {
	return &CarRepository{
		Type: Type,
		Speed: Speed,
	}
}

type CarService struct {
	CarRepository *CarRepository
}

func NewCarService(CarRepository *CarRepository) *CarService {
	return &CarService{
		CarRepository: CarRepository,
	}
}

