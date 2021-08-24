package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type configs struct {
	appPort                 string
	userServiceGRPCPort     string
	employeeServiceGRPCPort string
}

var config configs

func Init() (err error) {
	err = LoadConfig("./", "application")
	if err != nil {
		return
	}
	err = CheckIfSet("APP_PORT")
	if err != nil {
		return
	}
	appPort := viper.GetString("APP_PORT")
	config.appPort = appPort

	err = CheckIfSet("USER_GRPC_PORT")
	if err != nil {
		return
	}
	userServiceGRPCPort := viper.GetString("USER_GRPC_PORT")
	config.userServiceGRPCPort = userServiceGRPCPort

	err = CheckIfSet("EMPLOYEE_GRPC_PORT")
	if err != nil {
		return
	}
	employeeServiceGRPCPort := viper.GetString("EMPLOYEE_GRPC_PORT")
	config.employeeServiceGRPCPort = employeeServiceGRPCPort

	fmt.Println("--------------------------------------------")
	fmt.Printf("CONFIG:%+v", config)

	return
}

func LoadConfig(filePath, fileName string) (err error) {
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filePath)

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		return
	}
	return
}

func CheckIfSet(key string) (err error) {
	if !viper.IsSet(key) {
		err = errors.New(fmt.Sprintf("Key %s is not set", key))
		return
	}
	return
}

func GetAppPort() string {
	return config.appPort
}

func GetUserServiceGRPCPort() string {
	return config.userServiceGRPCPort
}

func GetEmployeeServiceGRPCPort() string {
	return config.employeeServiceGRPCPort
}
