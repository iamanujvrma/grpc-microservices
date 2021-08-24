package main

import (
	"fmt"
	"net/http"

	"github.com/iamanujvrma/grpc-microservices/mobile_core_service/config"
	"github.com/iamanujvrma/grpc-microservices/mobile_core_service/server"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("error in init. error: %s", err.Error())
	}

	err = http.ListenAndServe(":"+config.GetAppPort(), server.Routes())
	if err != nil {
		fmt.Printf("error in listen and serve. error: %s", err.Error())
	}
}
