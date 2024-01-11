package pkg

import (
	"fuux/internal/entity"
	"github.com/spf13/viper"
	"log"
)

//func Config(f *entity.Flag) (*entity.Config, error) {
//	viper.SetConfigType("yml")
//	viper.SetConfigName(*f.Config)
//	viper.AddConfigPath("./config")
//
//	if err := viper.ReadInConfig(); err != nil {
//		return nil, err
//	}
//
//	var config entity.Config
//	if err := viper.Unmarshal(&config); err != nil {
//		return nil, err
//	}
//
//	return &config, nil
//}

func NewConfig(env string) (*entity.Config, error) {
	var config *entity.Config

	viper.SetConfigType("yml")
	viper.SetConfigName(env)
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	return config, nil
}
