package main

import (
	"fmt"
	"google_tts_wrapper/routers"
)

func init() {
	fmt.Println("Initializing App")
	routers.InitializeRouters()
}
func main() {
	routers.Router.Run()
}
