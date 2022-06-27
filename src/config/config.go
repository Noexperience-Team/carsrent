package config

import (
	"path"
	"runtime"
	"time"

	"github.com/jinzhu/configor"
)

type Parameters struct {
	Env string `default:"dev" env:"CONFIGOR_ENV"`
}

type Server struct {
	Type           string
	Host           string        `default:"localhost" env:"SERVER_HOST"`
	Port           string        `default:"8888" env:"SERVER_PORT"`
	ReadTimeout    time.Duration `default:"1000" env:"READ_TIMEOUT"`
	WriteTimeout   time.Duration `default:"1000" env:"WRITE_TIMEOUT"`
	MaxHeaderBytes int           `default:"1000" env:"MAX_HEADER_BYTES"`
}

type Config struct {
	Parameters Parameters `yaml:"parameters"`
	Server     Server     `yaml:"server"`
}

// MakeConfig sets the application config and change configuration according to env variable
// returns config,error
func MakeConfig() (conf *Config, err error) {
	var configFilePath string
	config := configor.New(&configor.Config{})
	switch config.GetEnvironment() {
	case "dev":
		configFilePath = "../../config/config.dev.yml"
	case "prod":
		configFilePath = "../../config/config.prod.yml"
	default:
		configFilePath = "../../config/config.dev.yml"
	}
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), configFilePath)
	// Load configuration from yml file
	conf = new(Config)
	err = config.Load(conf, filepath)
	return
}
