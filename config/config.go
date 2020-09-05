package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var mu sync.Mutex

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	LoadApp()
	LoadDB()

}
