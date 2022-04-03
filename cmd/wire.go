//go:build wireinject
// +build wireinject

package main

import (
	"gin-based-microservice-demo/internal/application"
	"github.com/google/wire"
)

func initializeApplication(name application.Name) (*application.Application, error) {
	wire.Build(wireSet)
	return nil, nil
}
