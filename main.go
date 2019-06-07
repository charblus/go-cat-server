package main

import (
	"fmt"
	"go-cat-server/service"
)

var appName = "catService"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}