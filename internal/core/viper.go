package core

import (
	"log"

	"github.com/spf13/viper"
)

var defaultConfig = "./config/config.yaml"

func InitConfig(configFile string) (err error) {

	if configFile == "" {
		configFile = defaultConfig
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed, err: %v", err)
		return
	}

	if err = viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("unmarshal config failed, err: %v", err)
		return
	}

	return
}
