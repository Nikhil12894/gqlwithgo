package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	DBConfig struct {
		Port   string `json:"port"`
		DbName string `json:"dbname"`
		Vender string `json:"vender"`
	} `json:"dbconfig"`
	Server struct {
		Port string `json:"port"`
		Host string `json:"host"`
	} `json:"serverconfig"`
	IsProd bool `json:"isprod"`
}

func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
