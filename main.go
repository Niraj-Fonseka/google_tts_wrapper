package main

import (
	"fmt"
	"google_tts_wrapper/models"
	"google_tts_wrapper/routers"
)

func init() {
	fmt.Println("Initializing App")
	models.InitializeDatastore()
	routers.InitializeRouters()
}
func main() {
	routers.Router.Run()
}
