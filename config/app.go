package config

import (
	"time"

	"github.com/spf13/viper"
)

// AppConfig ...
type AppConfig struct {
	Port         int
	Name         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Debug        bool
}

var appCfg AppConfig

// LoadApp ...
func LoadApp() {
	appCfg = AppConfig{
		Port:         viper.GetInt("app.port"),
		Name:         viper.GetString("app.name"),
		ReadTimeout:  viper.GetDuration("app.read_timeout") * time.Second,
		WriteTimeout: viper.GetDuration("app.write_timeout") * time.Second,
		IdleTimeout:  viper.GetDuration("app.idle_timeout") * time.Second,
	}
}

// App ...
func App() AppConfig {
	return appCfg
}
