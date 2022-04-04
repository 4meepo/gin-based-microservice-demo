package main

import (
	"gin-based-microservice-demo/internal/application"
	"github.com/google/wire"
)

func main() {
	app, err := initializeApplication("gin-based-microservice-demo")

	if nil != err {
		panic(err)
	}
	err = app.Run()
	if err != nil {
		panic(err)
	}
}

var wireSet = wire.NewSet(
	application.New,
)
