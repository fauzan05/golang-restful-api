package simple

type Name string
type Age int

type OwnerRepository struct {
	Name Name
	Age Age
}

func NewOwnerRepository(Name Name, Age Age) *OwnerRepository {
	return &OwnerRepository{
		Name: Name,
		Age: Age,
	}
}

type OwnerService struct {
	OwnerRepository *OwnerRepository
}

func NewOwnerService(OwnerRepository *OwnerRepository) *OwnerService {
	return &OwnerService{
		OwnerRepository: OwnerRepository,
	}
}