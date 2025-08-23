package core

type GoUnoConfig struct {
	ApiServerConfig ApiServerConfig `mapstructure:"api_server"`
}

type ApiServerConfig struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
	Debug   bool   `mapstructure:"debug"`
}

var GlobalConfig GoUnoConfig
