module github.com/iamanujvrma/grpc-microservices/mobile_core_service

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/iamanujvrma/grpc-microservices/common v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.8.1
	google.golang.org/grpc v1.40.0
)

replace github.com/iamanujvrma/grpc-microservices/common => ../common
