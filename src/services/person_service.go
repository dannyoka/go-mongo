package services

import (
	"github.com/dannyoka/go-mongo/src/data"
	"github.com/dannyoka/go-mongo/src/repositories"
)

type IPersonService interface {
	FindOneByName(name string) (data.Person, error)
}

type PersonService struct {
	repository repositories.IPersonRepository
}

func NewPersonService(repository repositories.IPersonRepository) IPersonService {
	return &PersonService{
		repository: repository,
	}
}

func (s *PersonService) FindOneByName(name string) (data.Person, error) {
	return s.repository.FindOneByName(name)
}
