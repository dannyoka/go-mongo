package services

import "github.com/dannyoka/go-mongo/src/repositories"

type ServicesContext struct {
	PersonService IPersonService
}

func InitServices(repositoriesContext repositories.RepositoriesContext) ServicesContext {
	personService := NewPersonService(repositoriesContext.PersonRepository)
	servicesContext := ServicesContext{
		PersonService: personService,
	}
	return servicesContext
}
