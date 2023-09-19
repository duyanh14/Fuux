package entity

type Config struct {
	Env      string         `mapstructure:"env"`
	Listen   string         `mapstructure:"listen"`
	Database ConfigDatabase `mapstructure:"database"`
	JWT      ConfigJWT      `mapstructure:"jwt"`
}

type ConfigDatabase struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Token    string `mapstructure:"token"`
}

type ConfigJWT struct {
	Secret string `mapstructure:"secret"`
}
