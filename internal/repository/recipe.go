package repository

import "gin-based-microservice-demo/internal/model"

type RecipeRepository interface {
	FindByID(id string) (*model.Recipe, error)
	Save(recipe *model.Recipe) error
}

type RecipeRepositoryImpl struct{}

func (RecipeRepositoryImpl) FindByID(id string) (*model.Recipe, error) {
	// TODO implement me
	panic("implement me")
}

func (RecipeRepositoryImpl) Save(recipe *model.Recipe) error {
	// TODO implement me
	panic("implement me")
}
