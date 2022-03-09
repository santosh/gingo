package main

import "github.com/santosh/gingo/routes"

func main() {
	router := routes.SetupRouter()

	router.Run(":8080")
}

