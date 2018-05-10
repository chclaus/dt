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
	Hash   HashConfig
}

type ServerConfig struct {
	Address string
	Port    string
}

type Base64Config struct {
	Encoding string
}

type RandomConfig struct {
	Algorithm string
	Length    int
}

type UUIDConfig struct {
	Namespace string
	Version   int
}

type HashConfig struct {
	Algorithm string
	Cost      int
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
