package main

import (
	"fmt"
	_ "kiva/cmd/app"
	"kiva/settings"
)

func main() {
	fmt.Println("Hello I am start of Kiva Project")
	settings.ReadSettings(`./settings.json`)
}
