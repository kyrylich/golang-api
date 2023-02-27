package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

const securityKey = "security"
const signingKeyKey = "signingKey"

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

type SecurityConfig struct {
	SigningKey string `mapstructure:"signingKey"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Security SecurityConfig `mapstructure:"security"`
}

func InitConfiguration() (*Config, error) {
	_, filename, _, _ := runtime.Caller(0) // get path to current file: config.go. Maybe it's a hack, but it works
	configPath := path.Join(path.Dir(filename), "../../config")

	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	err := viper.Unmarshal(&config)

	return &config, err
}

func GetSigningKey() []byte {
	return []byte(viper.GetStringMapString(securityKey)[signingKeyKey])
}
