package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server ServerConfig
	Base64 Base64Config
	Random RandomConfig
	UUID   UUIDConfig
}

type ServerConfig struct {
	Address string
	Port    string
}

type Base64Config struct {
	Encoding string
}

type RandomConfig struct {
	Generator string
	Length    uint32
	Cost      uint8
}

type UUIDConfig struct {
	Namespace string
}

var Cfg *Config

func InitConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigName(".dt")
	viper.AddConfigPath(home)
	viper.ReadInConfig()
	viper.Unmarshal(&Cfg)
}
