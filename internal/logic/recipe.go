package logic

import (
	"context"
	"gin-based-microservice-demo/internal/model"

	"gin-based-microservice-demo/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeLogic interface {
	GetRecipeByUuid(ctx context.Context, uuid string) (*model.Recipe, error)
}

type RecipeLogicImpl struct {
	repository.RecipeRepository
}

func NewRecipeLogicImpl(recipeRepository repository.RecipeRepository) (*RecipeLogicImpl, error) {
	return &RecipeLogicImpl{recipeRepository}, nil
}

func (r *RecipeLogicImpl) GetRecipeByUuid(ctx context.Context, uuid string) (*model.Recipe, error) {
	return &model.Recipe{
		ID:   primitive.NewObjectID(),
		Name: "recipe1",
	}, nil
}
