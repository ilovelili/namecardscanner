package main

import "namecardscanner/app"

// main Entry
func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
