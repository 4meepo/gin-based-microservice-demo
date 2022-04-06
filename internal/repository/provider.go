package repository

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewRecipeRepositoryImpl,
	wire.Bind(new(RecipeRepository), new(*RecipeRepositoryImpl)),
)
