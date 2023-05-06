package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env string

const (
	EnvProd Env = "PRODUCTION"
	EnvDev  Env = "DEVELOPMENT"
)

type Config struct {
	AppHost        string   `mapstructure:"APP_HOST"`
	AppPort        string   `mapstructure:"APP_PORT"`
	AppDomain      string   `mapstructure:"APP_DOMAIN"`
	PostgresDSN    string   `mapstructure:"DB_CONNECTION_URL"`
	Env            Env      `mapstructure:"ENV"`
	AllowedOrigins []string `mapstructure:"ORIGINS"`
}

func New(filePath string) Config {
	var config Config

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}

	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("APP_PORT", "8000")
	viper.SetDefault("ENV", EnvDev)
	viper.SetDefault("ORIGINS", [1]string{"http://localhost:3000"})
	viper.Unmarshal(&config)

	if config.Env != EnvDev && config.Env != EnvProd {
		config.Env = EnvProd
	}

	return config
}
