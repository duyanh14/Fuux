package usecase

import (
	"encoding/json"
	"fuux/internal/entity"
	"io"
	"os"
)

type config struct {
}

var Config *config

func NewConfig(flag *entity.Flag) (*config, error) {
	Config = &config{}

	Config.load(flag)

	return Config, nil
}

func (uc *config) load(flag *entity.Flag) (*entity.Config, error) {
	file, err := os.Open(*flag.Config)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cf *entity.Config

	if err := json.Unmarshal(content, &cf); err != nil {
		return nil, err
	}

	return cf, nil
}
