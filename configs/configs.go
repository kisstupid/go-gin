package configs

import (
	"log"

	"github.com/spf13/viper"
)

var configs *viper.Viper

func Init(env string) {
	var err error

	configs = viper.New()
	configs.SetConfigFile("yaml")
	configs.SetConfigName(env)
	configs.AddConfigPath("configs/")

	err = configs.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return configs
}
