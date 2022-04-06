// +build wireinject

// The build tag makes sure the stub is not built in the final build.


package main

import (
	"gin-based-microservice-demo/internal/controller"
	"gin-based-microservice-demo/internal/logic"

	"gin-based-microservice-demo/internal/repository"
	appp "gin-based-microservice-demo/pkg/app"
	"gin-based-microservice-demo/pkg/protocal/http"

	"github.com/google/wire"
)

var wireSet = wire.NewSet(
	appp.ProviderSet,
	http.ProviderSet,
	controller.ProviderSet,
	logic.ProviderSet,
	repository.ProviderSet,

)
func initializeApplication(name appp.Name) (*appp.Application, error) {
	wire.Build(wireSet)
	return nil, nil
}
