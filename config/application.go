package config

import "github.com/SarkiMudboy/meeet/pkg/env"

type Config struct {
	Addr    string
	DB      DBConfig
	Storage *StorageConfig
}

func (c *Config) ServerAddr() string {
	return c.Addr
}

func LoadAppConfig() (Config, error) {

	cfg := Config{
		Addr:    env.GetString("SERVER_ADDR", ":8080"),
		DB:      *loadDBConfig(),
		Storage: NewStorageConfig("assets"),
	}

	return cfg, nil
}
