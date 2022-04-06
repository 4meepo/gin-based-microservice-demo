package main

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
