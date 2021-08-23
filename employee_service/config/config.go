package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type configs struct {
	grpcPort string
}

var config configs

func Init() (err error) {
	err = LoadConfig("./", "application")
	if err != nil {
		return
	}
	err = CheckIfSet("GRPC_PORT")
	if err != nil {
		return
	}
	grpcPort := viper.GetString("GRPC_PORT")
	config.grpcPort = grpcPort

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

func GetGRPCPort() string {
	return config.grpcPort
}
