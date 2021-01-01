package main

import (
	"log"

	"github.com/spf13/viper"
)

type MySQL struct {
	Enable bool
	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
}

type Redis struct {
	Enable bool
	DBHost string
	DBPort int
	DBPass string
	DBName int
}

type Config struct {
	Redis  Redis
	MySQL  MySQL
	Host   string
	Token  []string
	Listen string
}

func GetConfig() Config {
	var config Config

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err == nil {
		err = viper.Unmarshal(&config)
	}

	if err != nil {
		log.Fatalf("%s.\n", err)
	}

	if !config.MySQL.Enable && !config.Redis.Enable {
		log.Fatalf("do not have db.\n")
	}
	return config
}
