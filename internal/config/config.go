package config

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type AppConfig struct {
	Title  string `yaml:"title"`
	Server Server `yaml:"server"`
	Logger Logger `yaml:"logger"`
}

type Server struct {
	Bind string `yaml:"bind"`
}

type Logger struct {
	Level string `yaml:"level"`
}

func Default() *AppConfig {
	return &AppConfig{
		Title: "shortify",
		Server: Server{
			Bind: ":8080",
		},
		Logger: Logger{
			Level: "debug",
		},
	}
}

func Init(confPath string) (c *AppConfig, e error) {
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, errors.New("using default config cause err reading config-file")
	}

	c = Default()
	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		return nil, errors.New("using default config cause unmarshaling config-file error")
	}

	return c, nil
}
