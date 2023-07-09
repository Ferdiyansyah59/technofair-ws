package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Port int
		AppID string
	}
	Database struct {
		User string
		Password string
		Host string
		Port int
		Dbname string
	}
}

var Config *config

func InitConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("error read config", err)
	}

	cfg := config{}
	err = viper.Unmarshal(&cfg)

	if err != nil {
		fmt.Println("Error unmarshal config ", err)
	}

	Config = &cfg
}