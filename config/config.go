package config

import (
	"github.com/BurntSushi/toml"
)

var (
	conf *Config
)

type Config struct {
	App *struct{}
	Jop *struct {
		AppKey    string
		SecretKey string
		SiteId    string
	}
}

func GetConfig(file string) *Config {
	if nil == conf {
		if _, err := toml.DecodeFile(file, &conf); nil != err {
			panic(err)
		}
	}
	return conf
}
