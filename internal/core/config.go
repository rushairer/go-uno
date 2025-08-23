package core

type GoUnoConfig struct {
	WebServerConfig WebServerConfig `mapstructure:"web_server"`
}

type WebServerConfig struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
	Debug   bool   `mapstructure:"debug"`
}

var GlobalConfig GoUnoConfig
