package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func LoadFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type Config struct {
	ExoscaleConfig `yaml:"exoscale_config"`
}

type ExoscaleConfig struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}
