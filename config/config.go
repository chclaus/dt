package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

// Config is the root collection of command configurations
type Config struct {
	Server ServerConfig
	Base64 Base64Config
	Random RandomConfig
	UUID   UUIDConfig
	Hash   HashConfig
	JWT    JWTConfig
}

// ServerConfig allows configuration settings of the server cmd
type ServerConfig struct {
	Address     string
	Port        string
	OpenBrowser bool
}

// Base64Config allows configuration settings of the base64 cmd
type Base64Config struct {
	Encoding string
}

// RandomConfig allows configuration settings of the random cmd
type RandomConfig struct {
	Algorithm string
	Length    int
}

// UUIDConfig allows configuration settings of the uuid cmd
type UUIDConfig struct {
	Namespace string
	Version   int
}

// HashConfig allows configuration settings of the hash cmd
type HashConfig struct {
	Algorithm string
	Cost      int
}

// JWTConfig allows configuration settings of the jwt cmd
type JWTConfig struct {
	Secret       string
	SecretFile   string
	Base64Secret bool
}

// Cfg the root object of the configuration
var Cfg *Config

// InitConfig initializes the configuration if provided.
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
