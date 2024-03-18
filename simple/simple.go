package simple

import "errors"

type SimpleRepository struct {
	Error bool
	Word string
}

func NewSimpleRepository(isError bool, word string) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
		Word: word,
	}
}

type SimpleService struct {
	SimpleRepository *SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("gagal membuat objek service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
