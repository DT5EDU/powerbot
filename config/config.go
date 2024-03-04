package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PowerApi          string
	FeishuAppId       string
	FeishuAppSecret   string
	VerificationToken string
	EventKey          string
	Database          string
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	return &Config{
		PowerApi:          viper.GetString("powerApi"),
		FeishuAppId:       viper.GetString("feishuAppId"),
		FeishuAppSecret:   viper.GetString("feishuAppSecret"),
		VerificationToken: viper.GetString("verificationToken"),
		EventKey:          viper.GetString("eventKey"),
		Database:          viper.GetString("database"),
	}
}
