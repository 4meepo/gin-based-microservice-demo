package controller

import (
	"gin-based-microservice-demo/internal/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecipeController struct {
	L logic.RecipeLogic
}

func NewRecipeController(l logic.RecipeLogic) (*RecipeController, error) {
	return &RecipeController{l}, nil
}

func (rc *RecipeController) GetRecipeByUuid(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	recipe, err := rc.L.GetRecipeByUuid(ctx, uuid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, recipe)
}
