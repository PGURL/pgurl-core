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
	environment := os.Getenv("ENVIRONMENT")
	log.Printf("%v\n", environment)

	v := viper.New()
	if (environment != "PRODUCTION") {
		v.SetConfigName("development")
	} else {
		v.SetConfigName("production")
	}

	v.SetConfigType("json")
	v.AddConfigPath("confs/")
	err_readconfig := v.ReadInConfig()

	if err_readconfig != nil {
		log.Fatalf("No Config: \n%#v", v)
	} else {
		config = v
	}
}

func GetConfig() *viper.Viper {
	return config
}

