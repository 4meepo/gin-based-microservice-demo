package controller

import (
	"gin-based-microservice-demo/pkg/protocal/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProviderSet is a wire provider set
var ProviderSet = wire.NewSet(
	Register,

	NewRecipeController,
)

func Register(recipeController *RecipeController) http.RegisterController {
	return func(r *gin.Engine) {
		r.GET("/recipe/:uuid", recipeController.GetRecipeByUuid)
	}
}
