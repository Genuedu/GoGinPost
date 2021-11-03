package utils

import "github.com/spf13/viper"

type Config struct {
	AppPort       string `mapstructrure:"APP_PORT"`
	RuntimeSetup  string `mapstructrure:"RUNTIME_SETUP"`
	DBDriver      string `mapstructrure:"DB_DRIVER"`
	ServerAddress string `mapstructrure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	err = viper.Unmarshal(&config)
	return
}
