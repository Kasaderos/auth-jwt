package config

import (
	"github.com/spf13/viper"
)

type ConfSt struct {
	Debug      bool   `mapstructure:"DEBUG"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	PgDsn      string `mapstructure:"PG_DSN"`
	HttpListen string `mapstructure:"HTTP_LISTEN"`
	HttpCors   bool   `mapstructure:"HTTP_CORS"`
	CertPath   string `mapstructure:"CERT_PATH"`
	CertPsw    string `mapstructure:"CERT_PSW"`
	BaseURL    string `mapstructure:"BASE_URL"`
}

func Load() *ConfSt {
	result := &ConfSt{}

	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_LISTEN", ":8888")
	viper.SetDefault("BASE_URL", "http://localhost:8888/")
	viper.SetDefault("CERT_PATH", "")
	viper.SetDefault("CERT_PSW", "")

	viper.AutomaticEnv()

	result.BotToken = viper.GetString("BASE_URL")
	result.PgDsn = viper.GetString("PG_DSN")
	result.HttpListen = viper.GetString("HTTP_LISTEN")
	result.LogLevel = viper.GetString("LOG_LEVEL")
	result.HttpCors = viper.GetBool("HTTP_CORS")
	result.CertPath = viper.GetString("CERT_PATH")
	result.CertPsw = viper.GetString("CERT_PSW")

	return result
}
