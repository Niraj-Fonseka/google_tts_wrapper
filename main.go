package main

import (
	"fmt"
	"webapp_go/models"
	"webapp_go/routers"
)

func init() {
	fmt.Println("Initializing App")
	models.InitializeDatastore()
	routers.InitializeRouters()
}
func main() {
	routers.Router.Run()
}
