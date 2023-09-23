package pkg

import (
	"fuux/internal/entity"
	"github.com/spf13/viper"
)

func Config(f *entity.Flag) (*entity.Config, error) {
	viper.SetConfigType("yml")
	viper.SetConfigName(*f.Config)
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config entity.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
