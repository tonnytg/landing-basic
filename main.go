package main

import (
	"log"

	"github.com/tonnytg/lading-basic/pkg/webserver"
)

func main() {
	log.Println("start webserver")

	webserver.Start()
}
