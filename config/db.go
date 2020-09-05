package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// DatabaseConfig holds the database configuration
type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	Debug           bool
}

var dbCfg DatabaseConfig

// DB returns the default database configuration
func DB() DatabaseConfig {
	return dbCfg
}

// LoadDB loads database configuration
func LoadDB() {
	mu.Lock()
	defer mu.Unlock()
	dbCfg = DatabaseConfig{
		Name:            viper.GetString("database.name"),
		User:            viper.GetString("database.username"),
		Password:        viper.GetString("database.password"),
		Host:            fmt.Sprintf("%s:%d", viper.GetString("database.host"), viper.GetInt("database.port")),
		Port:            viper.GetInt("database.port"),
		MaxIdleConn:     viper.GetInt("database.max_idle_connection"),
		MaxOpenConn:     viper.GetInt("database.max_open_connection"),
		MaxConnLifetime: viper.GetDuration("database.max_connection_lifetime") * time.Second,
		Debug:           viper.GetBool("database.debug"),
	}
}
