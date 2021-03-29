package main

import (
	"fmt"
	"kiva/cmd/app"

	//_ "kiva/cmd/app"
	"kiva/settings"
)

func main() {
	fmt.Println("Hello I am start of Kiva Project")
	settings.ReadSettings(`./settings.json`)
	fmt.Println(settings.ReqURL)
//	fmt.Println(readSettings)
	app.InitRoutes()

}
