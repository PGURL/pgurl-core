package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configuration struct {
	Port int
}

var config *viper.Viper

func init() {
	environment := os.Getenv("ENV_MODE")
	log.Printf("ENV_MODE: %v\n", environment)

	v := viper.New()
	if environment != "PRODUCTION" {
		environment = "development"
	} else {
		environment = "production"
	}
	v.SetConfigName(environment)

	v.SetConfigType("json")
	v.AddConfigPath("/confs")
	v.AddConfigPath("../../confs/")
	v.AddConfigPath("confs/")
	err_readconfig := v.ReadInConfig()

	if err_readconfig == nil {
		config = v
		log.Printf("Now use conf %v", environment)
	} else {
		log.Fatalf("No Config: \n%#v \n%v", v, err_readconfig)
	}
}

func GetConfig() *viper.Viper {
	return config
}
