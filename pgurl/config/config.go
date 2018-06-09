package config

import (
	"os"
	"log"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port int
}

var config *viper.Viper

func init() {
	environment := os.Getenv("ENV_MODE")
	log.Printf("Load conf: %v\n", environment)

	v := viper.New()
	if (environment != "PRODUCTION") {
		v.SetConfigName("development")
	} else {
		v.SetConfigName("production")
	}

	v.SetConfigType("json")
	v.AddConfigPath("/")
	v.AddConfigPath("../../confs/")
	v.AddConfigPath("confs/")
	err_readconfig := v.ReadInConfig()

	if err_readconfig == nil {
		config = v
	} else {
		log.Fatalf("No Config: \n%#v \n%v", v, err_readconfig)
	}
}

func GetConfig() *viper.Viper {
	return config
}

