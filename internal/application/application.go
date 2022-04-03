package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Name is the unique name of this app
type Name string

type Application struct {
	Name Name
	*gin.Engine
	// TODO: grpc server??
}

func NewApplication(name Name) *Application {
	return &Application{
		Name:   name,
		Engine: gin.Default(),
	}
}

func (app *Application) Run() error {
	err := app.Engine.Run()
	if nil != err {
		return err
	}

	// FIXME: gin.Run 会阻塞 main goroutine
	fmt.Println("start grpc ???")
	return nil
}
