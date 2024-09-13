package main

import (
	"github.com/GabrielPereira187/go-dynamo/initializers"
	"github.com/GabrielPereira187/go-dynamo/router"
)

func init() {
	initializers.LoadDotEnv()
	router.Initialize()
}

func main() {
}
