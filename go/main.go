package main

import (
	"github.com/Khmer495/social-creators-teamG-server/go/infrastracture"
)

func main() {
	e := infrastracture.ApiServer()
	e.Logger.Fatal(e.Start(":8080"))
}
