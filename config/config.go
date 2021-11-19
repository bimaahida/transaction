package config

import (
	"github.com/spf13/viper"
)

func GetConfig() (Configuration, error) {
	var conf Configuration

	viper.SetConfigFile(`config/config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return Configuration{}, err
	}

	return conf, nil
}
