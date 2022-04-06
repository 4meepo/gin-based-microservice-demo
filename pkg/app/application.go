package application

import (
	"fmt"
	"gin-based-microservice-demo/pkg/protocal/http"
)

// Name is the unique name of this app
type Name string

type Application struct {
	*http.Server
	// TODO: grpc.Server
	Name Name
}

func NewApplication(name Name, httpServer *http.Server) *Application {
	return &Application{
		httpServer,
		name,
	}
}

func (app *Application) Startup() error {
	err := app.Run()
	if err != nil {
		panic(err)
	}

	// FIXME: gin.Run 会阻塞 main goroutine
	fmt.Println("start grpc ???")
	return nil
}
