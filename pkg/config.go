package pkg

import (
	"fuux/internal/entity"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Config(f *entity.Flag) *entity.Config {
	viper.SetConfigType("yml")
	viper.SetConfigName(*f.Config)
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	var config entity.Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
