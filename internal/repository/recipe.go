package repository

import "gin-based-microservice-demo/internal/model"

type RecipeRepository interface {
	FindByID(id string) (*model.Recipe, error)
	Save(recipe *model.Recipe) error
}

type RecipeRepositoryImpl struct{}

func NewRecipeRepositoryImpl() *RecipeRepositoryImpl{
	return &RecipeRepositoryImpl{}
}

func (r *RecipeRepositoryImpl) FindByID(id string) (*model.Recipe, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RecipeRepositoryImpl) Save(recipe *model.Recipe) error {
	panic("not implemented") // TODO: Implement
}
