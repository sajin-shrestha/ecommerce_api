package main

import "github.com/sajin-shrestha/ecommerce_api/api"

func main() {
	apiServer := api.NewAPIServer(":9999")
	if err := apiServer.Run(); err != nil {
		panic(err.Error())
	}
}
