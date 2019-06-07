package main

import (
	"fmt"
	"github.com/go-cat/service"
)

var appName = "catService"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}