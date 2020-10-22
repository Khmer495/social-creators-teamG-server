package main

import (
	"github.com/Khmer495/social-creators-teamG-server/go/infrastracture"
)

func main() {
	infrastracture.InitDBServer()
	infrastracture.InitApiServer()
}
