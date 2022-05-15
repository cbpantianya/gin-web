package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Server struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	} `toml:"Server"`
	Database struct {
		MySql struct {
			Host     string `toml:"host"`
			Port     string `toml:"port"`
			User     string `toml:"user"`
			Password string `toml:"password"`
			Database string `toml:"database"`
		} `toml:"MySql"`
		Redis struct {
			Host     string `toml:"host"`
			Port     string `toml:"port"`
			Password string `toml:"password"`
			Database int    `toml:"database"`
		} `toml:"Redis"`
	} `toml:"DataBase"`
}

var Cfg Config

func InitConfig() {
	_, err := toml.DecodeFile("./server/config/config.toml", &Cfg)
	if err != nil {
		logrus.WithField("server", "Config").Info("Load config failed!")
	}
	return
}
