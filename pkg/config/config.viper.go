package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type viperConfig struct {
	Server `mapstructure:",squash"`
	Db     `mapstructure:",squash"`
	Jwt    `mapstructure:",squash"`
	Aws    `mapstructure:",squash"`
}

var (
	once     sync.Once
	instance Config
)

func NewViperConfig() Config {
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile(".env")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("Error reading configs file: %s", err)
		}

		cfg := &viperConfig{}

		err := v.Unmarshal(cfg)
		if err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}

		instance = cfg
	})

	return instance
}

func GetConfig() Config {
	if instance == nil {
		instance = NewViperConfig()
	}
	return instance
}

func (c *viperConfig) GetServer() Server {
	return c.Server
}

func (c *viperConfig) GetDb() Db {
	return c.Db
}

func (c *viperConfig) GetJwt() Jwt {
	return c.Jwt
}

func (c *viperConfig) GetAws() Aws {
	return c.Aws
}
