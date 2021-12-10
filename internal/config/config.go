package config

import (
	"errors"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type AppConfig struct {
	Title  string `yaml:"title"`
	Server Server `yaml:"server"`
	Logger Logger `yaml:"logger"`
	DB     DB     `yaml:"postgres"`
}

type Server struct {
	Bind string `yaml:"bind"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SslMode  string `yaml:"sslmode"`
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
		DB: DB{
			Host:     "localhost",
			Port:     "8082",
			User:     "root",
			Password: "root",
			DBName:   "postgres_db",
			SslMode:  "disable",
		},
	}
}

func Init(confPath string) (c *AppConfig, e error) {
	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		return nil, errors.New("using default config cause err reading config-file")
	}

	// TODO think about DEFAULT
	c = Default()
	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		return nil, errors.New("using default config cause unmarshalling config-file error")
	}

	return c, nil
}
