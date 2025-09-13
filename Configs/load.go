package Configs

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	App struct {
		Name        string `map structure:"name"`
		Port        int    `map structure:"port"`
		Environment string `map structure:"environment"`
	} `map structure:"app"`

	Database struct {
		Host     string `map structure:"host"`
		Port     int    `map structure:"port"`
		DBName   string `map structure:"dbname"`
		User     string `map structure:"user"`
		Password string `map structure:"password"`
	} `map structure:"database"`

	Logging struct {
		Level string `map structure:"level"`
		File  string `map structure:"file"`
	} `map structure:"logging"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set defaults for sensitive values
	viper.SetDefault("database.sslmode", "disable")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
