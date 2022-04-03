package service

import "gin-based-microservice-demo/internal/model"

type RecipeService interface {
	// GetRecipe GetAll returns all recipes
	GetRecipe(id string) (model.Recipe, error)
	// CreateRecipe GetRecipes GetAll returns all recipes
	CreateRecipe(recipe model.Recipe) (model.Recipe, error)
}
