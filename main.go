package main

import "github.com/sajin-shrestha/ecommerce_api/controllers"

func main() {
	apiServer := controllers.NewAPIServer(":9999")
	if err := apiServer.Run(); err != nil {
		panic(err)
	}
}
