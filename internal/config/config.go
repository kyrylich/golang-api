package config

import (
	"github.com/spf13/viper"
)

const securityKey = "security"
const signingKeyKey = "signingKey"

func InitConfiguration() error {
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func GetSigningKey() []byte {
	return []byte(viper.GetStringMapString(securityKey)[signingKeyKey])
}
