package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DevelopmentMode   bool   `mapstructure:"DEVELOPMENT_MODE"`
	HttpServerPort    string `mapstructure:"HTTP_SERVER_PORT"`
	HttpClientTimeout int    `mapstructure:"HTTP_CLIENT_TIMEOUT_MS"`
	CorreiosBaseURL   string `mapstructure:"CORREIOS_BASE_URL"`
	ViaCepBaseURL     string `mapstructure:"VIACEP_BASE_URL"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
