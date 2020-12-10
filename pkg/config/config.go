package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Env           string `mapstructure:"env"`
	Provider      string `mapstructure:"provider"`
	Region        string `mapstructure:"region"`
	Role          string `mapstructure:"role"`
	Elasticsearch struct {
		BootstrapServers []string `mapstructure:"bootstrap_servers"`
		IndexName        string   `mapstructure:"index_name"`
	} `mapstructure:"elasticsearch"`
}

//LoadConfig load the configuration from file
func LoadConfig() *Config {
	var c = &Config{}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Errorf("unable to parse configuration file. err=%v", err.Error())
		return &Config{}
	}
	return c
}
