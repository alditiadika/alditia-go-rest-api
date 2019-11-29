package main

import (
	"fmt"
	"github.com/alditiadika/alditia-go-rest-api/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	fmt.Println("HTTP REST run at port 3000")
	app.Run(":3000")
}
