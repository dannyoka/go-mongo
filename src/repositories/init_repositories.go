package repositories

import "go.mongodb.org/mongo-driver/mongo"

type RepositoriesContext struct {
	PersonRepository IPersonRepository
}

func InitRepositories(db *mongo.Database) RepositoriesContext {
	personRepository := NewPersonRepository(db)
	repositoriesContext := RepositoriesContext{
		PersonRepository: personRepository,
	}
	return repositoriesContext
}
